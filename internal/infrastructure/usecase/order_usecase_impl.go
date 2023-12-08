package usecase

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/ashkarax/vegn-eCommerce/pkg/razorpay"
	"github.com/go-playground/validator/v10"
	// requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
)

type OrderUseCase struct {
	OrderRepo      interfaceRepository.IOrderRepo
	CartRepo       interfaceRepository.ICartRepository
	DishRepo       interfaceRepository.IDishRepo
	AddressRepo    interfaceRepository.IAddressRepo
	UserRepo       interfaceRepository.IuserRepo
	CouponRepo     interfaceRepository.ICouponRepository
	RestaurantRepo interfaceRepository.IrestaurantRepo
	RazorKeys      *config.RazorPay
}

func NewOrderUseCase(orderRepo interfaceRepository.IOrderRepo,
	cartRepo interfaceRepository.ICartRepository,
	dishRepo interfaceRepository.IDishRepo,
	addressRepo interfaceRepository.IAddressRepo,
	userRepo interfaceRepository.IuserRepo,
	couponRepo interfaceRepository.ICouponRepository,
	restaurantRepo interfaceRepository.IrestaurantRepo,
	razorKeys *config.RazorPay) interfaceUseCase.IOrderUseCase {
	return &OrderUseCase{OrderRepo: orderRepo,
		CartRepo:       cartRepo,
		DishRepo:       dishRepo,
		AddressRepo:    addressRepo,
		RazorKeys:      razorKeys,
		UserRepo:       userRepo,
		RestaurantRepo: restaurantRepo,
		CouponRepo:     couponRepo}
}

func (r *OrderUseCase) PlaceNewOrder(orderDetails *requestmodels.OrderDetails) (*responsemodels.OrderDetailsRes, error) {
	var resOrder responsemodels.OrderDetailsRes
	totalPrice := 0.0
	finalPrice := 0.0
	var quantitySum uint
	var flag string
	var flagbool = 0
	var coupon *responsemodels.CouponDetails

	validate := validator.New(validator.WithRequiredStructEnabled())
	errValidate := validate.Struct(orderDetails)
	if errValidate != nil {
		if ve, ok := errValidate.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "AddressID":
					resOrder.AddressID = "should be a valid Address. "
				case "PaymentMethod":
					resOrder.PaymentMethod = "should be a valid Method. "
				case "CouponCode":
					resOrder.CouponCode = "should be a valid Code. "

				}
			}
			return &resOrder, errValidate
		}
	}

	addressVerifStat, erraddr := r.AddressRepo.VerifyAddress(orderDetails)
	if erraddr != nil {
		return &resOrder, erraddr
	}
	if addressVerifStat == 0 {
		return &resOrder, errors.New("address not matching with this user,enter a valid address")
	}

	cartDetailsSlice, err := r.CartRepo.FetchCartItemsofUser(&orderDetails.UserID)
	if err != nil {
		return &resOrder, err
	}

	for _, item := range *cartDetailsSlice {
		totalPrice += float64(item.Quantity) * item.Price
		quantitySum += item.Quantity
		if !item.Availability {
			flagbool = 1
			flag = fmt.Sprintf("%s dish in your cart is not available(dish id:%s) right now,remove it or add an available dish to continue", item.Name, item.DishID)
		}
		if item.Quantity > item.RemainingQuantity {
			flagbool = 2
			flag = fmt.Sprintf("%s dish in your cart is not have enough remaining quantity (dish id:%s) right now,remove it or try decrementing the quantity to continue", item.Name, item.DishID)
		}
	}
	if flagbool != 0 {
		if flagbool == 1 {
			return &resOrder, errors.New(flag)
		}
		if flagbool == 2 {
			return &resOrder, errors.New(flag)
		}
		return &resOrder, errors.New("flagbool value changed,error in program logic,check order-usecase-line82")
	}

	finalPrice = totalPrice
	resOrder.TotalAmount = finalPrice

	if orderDetails.CouponCode != "" {
		var errCoup error
		coupon, errCoup = r.CouponRepo.GetcouponDetailsByCode(&orderDetails.CouponCode)
		if errCoup != nil {
			return &resOrder, errCoup
		}

		// currentTime := time.Now()

		// fmt.Println("------------", coupon.StartDate)
		// fmt.Println("------------", coupon.EndDate)
		// fmt.Println("------------", currentTime)

		// if !(currentTime.After(coupon.StartDate) && currentTime.Before(coupon.EndDate)) {
		// 	fmt.Println("err occord")
		// 	return &resOrder, fmt.Errorf("coupon is not valid: it is outside the valid date range")
		// }

		if !(totalPrice >= coupon.MinAmount && totalPrice <= coupon.MaxAmount) {

			return &resOrder, fmt.Errorf("coupon is not valid: final price is not within the amount constraints")
		}

		finalPrice = finalPrice - (finalPrice * coupon.DiscountPercentage / 100)

		orderDetails.CouponId = coupon.CouponID
	}

	for _, cartItem := range *cartDetailsSlice {
		errDecrement := r.DishRepo.DecrementDishQuantity(&cartItem.DishID, &cartItem.Quantity)
		if errDecrement != nil {
			return &resOrder, errDecrement
		}
	}
	if orderDetails.PaymentMethod == "COD" {
		orderDetails.OrderStatus = "processing"
		orderDetails.PaymentStatus = "pending"
	}
	if orderDetails.PaymentMethod == "ONLINE" {
		orderDetails.OrderStatus = "pending"
		orderDetails.PaymentStatus = "pending"

		newval := math.Round(finalPrice*100) / 100
		RazorPayOId, errRaz := razorpay.RazorPayInitialize(&newval, &r.RazorKeys.KeyId, &r.RazorKeys.SecrectKey)
		if errRaz != nil {
			fmt.Println("-----------", errRaz)
			return &resOrder, errRaz
		}

		orderDetails.OrderIdRazorPay = *RazorPayOId
	}

	resOrder.DiscountedAmount = finalPrice

	orderId, errPlacingOrder := r.OrderRepo.PlaceNewOrder(orderDetails)
	if errPlacingOrder != nil {
		return &resOrder, errPlacingOrder
	}
	resOrder.OrderId = *orderId

	for _, cartItem := range *cartDetailsSlice {
		errPlacingOrder := r.OrderRepo.AddDishesToOrderedItems(orderId, &cartItem, orderDetails)
		if errPlacingOrder != nil {
			return &resOrder, errPlacingOrder
		}

		errCartClear := r.CartRepo.DeleteCartAfterOrder(&cartItem.CartID)
		if errCartClear != nil {
			return &resOrder, errCartClear
		}

	}

	return &resOrder, nil
}

func (r *OrderUseCase) GetAllOrders(userId *string) (*[]responsemodels.OrderDetailsResponse, error) {
	verResSlice, err := r.OrderRepo.GetAllOrdersByUser(userId)
	if err != nil {
		return verResSlice, err
	}

	return verResSlice, nil
}

func (r *OrderUseCase) AllOrdersForARestaurant(restaurantId *string) (*[]responsemodels.OrderDetailsResponse, error) {
	orderResSlice, err := r.OrderRepo.OrdersForRestaurantById(restaurantId)
	if err != nil {
		return orderResSlice, err
	}

	return orderResSlice, nil
}

func (r *OrderUseCase) CancelOrderById(orderDetails *requestmodels.CanOrRetReq) error {

	dataStruct, errChk := r.OrderRepo.ReturnOrderStats(orderDetails)
	if errChk != nil {
		return errChk
	}

	if dataStruct.OrderStatus == "delivered" || dataStruct.OrderStatus == "cancelled" {
		errdata := fmt.Sprintf("Already %s ,now you can't cancel this order", dataStruct.OrderStatus)
		return errors.New(errdata)
	}

	newStatus := "cancelled"
	errUpdt := r.OrderRepo.ChangeOrderStatus(&orderDetails.OrderedItemsID, &newStatus)
	if errUpdt != nil {
		return errUpdt
	}

	if dataStruct.OrderStatus != "outfordelivery" {
		dishIDStr := strconv.FormatUint(uint64(dataStruct.DishID), 10)
		errUpdateDishQuantity := r.DishRepo.IncrementDishQuantity(&dishIDStr, &dataStruct.OrderQuantity)
		if errUpdateDishQuantity != nil {
			return errUpdateDishQuantity
		}
	}

	if dataStruct.PaymentMethod != "COD" && dataStruct.PaymentStatus == "success" {
		refundAmount := dataStruct.DishPrice * float64(dataStruct.OrderQuantity)
		errUpdateWallet := r.UserRepo.AddMoneyToWallet(&orderDetails.UserID, &refundAmount)
		if errUpdateWallet != nil {
			return errUpdateWallet
		}
		newStatus := "refunded"
		errPStatUpdt := r.OrderRepo.UpdatePaymentStatus(&orderDetails.OrderedItemsID, &newStatus)
		if errPStatUpdt != nil {
			return errPStatUpdt
		}
	}

	return nil
}

func (r *OrderUseCase) ReturnOrderById(orderDetails *requestmodels.CanOrRetReq) error {

	dataStruct, errChk := r.OrderRepo.ReturnOrderStats(orderDetails)
	if errChk != nil {
		return errChk
	}

	if dataStruct.OrderStatus != "delivered" {
		errdata := fmt.Sprintf("cant return a %s order", dataStruct.OrderStatus)
		return errors.New(errdata)
	}

	currentTime := time.Now()
	timeDifference := currentTime.Sub(dataStruct.DeliverDate)
	minutesDifference := int(timeDifference.Minutes())

	if minutesDifference > 120 {
		return errors.New("time exceeded,you can only return a product only before 2 hours from it's time of delivery")
	}

	newStatus := "return"
	errUpdt := r.OrderRepo.ChangeOrderStatus(&orderDetails.OrderedItemsID, &newStatus)
	if errUpdt != nil {
		return errUpdt
	}

	if dataStruct.PaymentStatus == "success" {
		refundAmount := dataStruct.DishPrice * float64(dataStruct.OrderQuantity)
		errUpdateWallet := r.UserRepo.AddMoneyToWallet(&orderDetails.UserID, &refundAmount)
		if errUpdateWallet != nil {
			return errUpdateWallet
		}
		newStatus := "refunded"
		errPStatUpdt := r.OrderRepo.UpdatePaymentStatus(&orderDetails.OrderedItemsID, &newStatus)
		if errPStatUpdt != nil {
			return errPStatUpdt
		}
	}

	return nil
}

func (r *OrderUseCase) ChangeStatusToPreparing(restaurantIdString *string, ordereditemsid *string) error {
	var newStatus string
	dataStruct, errChk := r.OrderRepo.RestReturnOrderStatus(restaurantIdString, ordereditemsid)
	if errChk != nil {
		return errChk
	}

	if dataStruct.OrderStatus != "processing" {
		errdata := fmt.Sprintf("cant change a %s status to preparing", dataStruct.OrderStatus)
		return errors.New(errdata)
	}
	newStatus = "preparing"
	errUpdt := r.OrderRepo.ChangeOrderStatus(ordereditemsid, &newStatus)
	if errUpdt != nil {
		return errUpdt
	}
	return nil

}

func (r *OrderUseCase) ChangeStatusToOutForDelivery(restaurantIdString *string, ordereditemsid *string) error {
	var newStatus string
	dataStruct, errChk := r.OrderRepo.RestReturnOrderStatus(restaurantIdString, ordereditemsid)
	if errChk != nil {
		return errChk
	}

	if dataStruct.OrderStatus != "preparing" {
		errdata := fmt.Sprintf("cant change a %s status to out-for-delivery", dataStruct.OrderStatus)
		return errors.New(errdata)
	}
	newStatus = "outfordelivery"
	errUpdt := r.OrderRepo.ChangeOrderStatus(ordereditemsid, &newStatus)
	if errUpdt != nil {
		return errUpdt
	}
	return nil

}

func (r *OrderUseCase) ChangeStatusToDelivered(restaurantIdString *string, ordereditemsid *string) error {
	var newStatus string
	dataStruct, errChk := r.OrderRepo.RestReturnOrderStatus(restaurantIdString, ordereditemsid)
	if errChk != nil {
		return errChk
	}

	if dataStruct.OrderStatus != "outfordelivery" {
		errdata := fmt.Sprintf("cant change a %s status to delivered", dataStruct.OrderStatus)
		return errors.New(errdata)
	}

	totalAmount := dataStruct.DishPrice * float64(dataStruct.OrderQuantity)

	if dataStruct.PaymentMethod == "COD" {
		err := r.RestaurantRepo.AddMoneyToCodWallet(&totalAmount, restaurantIdString)
		if err != nil {
			return err
		}
		newPaymentStat := "success"
		errChangeStat := r.OrderRepo.UpdatePaymentStatus(ordereditemsid, &newPaymentStat)
		if errChangeStat != nil {
			return errChangeStat
		}

	}
	if dataStruct.PaymentMethod == "ONLINE" || dataStruct.PaymentMethod == "WALLET" {
		err := r.RestaurantRepo.AddMoneyToAdminCredit(&totalAmount, restaurantIdString)
		if err != nil {
			return err
		}
	}
	newStatus = "delivered"
	errUpdt := r.OrderRepo.ChangeOrderStatus(ordereditemsid, &newStatus)
	if errUpdt != nil {
		return errUpdt
	}
	r.OrderRepo.UpdateDeliveryDate(ordereditemsid)
	return nil

}
