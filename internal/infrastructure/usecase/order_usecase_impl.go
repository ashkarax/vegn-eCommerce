package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ashkarax/vegn-eCommerce/internal/config"
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/ashkarax/vegn-eCommerce/pkg/aws"
	"github.com/ashkarax/vegn-eCommerce/pkg/razorpay"
	"github.com/go-playground/validator/v10"
	"github.com/jung-kurt/gofpdf"
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
	for i, dish := range *cartDetailsSlice {
		(*cartDetailsSlice)[i].SalePrice = math.Ceil(dish.Price - (dish.Price * float64(dish.DiscountPercentage) / 100))
	}

	for _, item := range *cartDetailsSlice {
		totalPrice += float64(item.Quantity) * item.SalePrice
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

		finalPrice = math.Ceil(finalPrice - (finalPrice * coupon.DiscountPercentage / 100))

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

		newval := math.Ceil((finalPrice * 100) / 100)
		RazorPayOId, errRaz := razorpay.RazorPayInitialize(&newval, &r.RazorKeys.KeyId, &r.RazorKeys.SecrectKey)
		if errRaz != nil {
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

// func (r *OrderUseCase) AllOrdersForARestaurant(restaurantId *string) (*[]responsemodels.OrderDetailsResponse, error) {
// 	orderResSlice, err := r.OrderRepo.OrdersForRestaurantById(restaurantId)
// 	if err != nil {
// 		return orderResSlice, err
// 	}

//		return orderResSlice, nil
//	}
func (r *OrderUseCase) AllOrdersForARestaurant(restaurantId *string) (*[]responsemodels.OrderResponseX, error) {
	var returnSlice []responsemodels.OrderResponseX

	var orderDetails responsemodels.OrderResponseX
	var orderedItemsDetails responsemodels.OrderDishDetailsX

	orderResSlice, err := r.OrderRepo.OrdersForRestaurantById(restaurantId)
	if err != nil {
		return &returnSlice, err
	}
	gatebool := true
	totalAmount := 0.00
	orderDetails.OrderID = (*orderResSlice)[0].OrderID

	for i := 0; i < len(*orderResSlice); i++ {

		if orderDetails.OrderID != (*orderResSlice)[i].OrderID {
			orderDetails.TotalAmount = totalAmount
			returnSlice = append(returnSlice, orderDetails)
			gatebool = true
			totalAmount = 0
			orderedItemsDetails = responsemodels.OrderDishDetailsX{}
			orderDetails = responsemodels.OrderResponseX{}
		}

		if gatebool {
			orderDetails.OrderID = (*orderResSlice)[i].OrderID
			gatebool = false
		}

		orderDetails.PaymentMethod = (*orderResSlice)[i].PaymentMethod
		orderDetails.PaymentStatus = (*orderResSlice)[i].PaymentStatus
		orderDetails.OrderDate = (*orderResSlice)[i].OrderDate
		orderDetails.FName = (*orderResSlice)[i].FName
		orderDetails.LName = (*orderResSlice)[i].LName
		orderDetails.Line1 = (*orderResSlice)[i].Line1
		orderDetails.Street = (*orderResSlice)[i].Street
		orderDetails.City = (*orderResSlice)[i].City
		orderDetails.PostalCode = (*orderResSlice)[i].PostalCode
		orderDetails.Phone = (*orderResSlice)[i].Phone
		orderDetails.AlternatePhone = (*orderResSlice)[i].AlternatePhone

		orderedItemsDetails.OrderedItemsID = (*orderResSlice)[i].OrderedItemsID
		orderedItemsDetails.DishID = (*orderResSlice)[i].DishID
		orderedItemsDetails.OrderID = (*orderResSlice)[i].OrderID
		orderedItemsDetails.OrderQuantity = (*orderResSlice)[i].OrderQuantity
		orderedItemsDetails.ImageURL1 = (*orderResSlice)[i].ImageURL1
		orderedItemsDetails.Name = (*orderResSlice)[i].Name
		orderedItemsDetails.OrderStatus = (*orderResSlice)[i].OrderStatus
		orderedItemsDetails.PortionSize = (*orderResSlice)[i].PortionSize
		orderedItemsDetails.DishPrice = (*orderResSlice)[i].DishPrice
		totalAmount += (*orderResSlice)[i].DishPrice

		orderDetails.OrderedItems = append(orderDetails.OrderedItems, orderedItemsDetails)

		if len(*orderResSlice)-1 == i {
			orderDetails.TotalAmount = totalAmount
			returnSlice = append(returnSlice, orderDetails)
		}
	}

	return &returnSlice, nil
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

func (r *OrderUseCase) GenerateInvoice(userId *string, orderId *string) (*string, error) {
	var pdfLink string
	var totalAmount float64
	var FinalAmount float64
	BucketFolder := "vegn-ecommerce-tempfiles/invoices/"

	orderDetails, errr := r.OrderRepo.GetOrderDataForPDFByIds(userId, orderId)
	if errr != nil {
		return &pdfLink, errr
	}
	orderedItemsDetails, err := r.OrderRepo.GetOrderDetailsForPDFById(orderId)
	if err != nil {
		return &pdfLink, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 10, "INVOICE")

	// Add platform details
	pdf.SetX(140)
	pdf.Cell(40, 10, " Veg*n")
	pdf.Ln(-1)
	pdf.SetX(140)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 6, " Online Veg Food Delivery")
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 6, fmt.Sprintf("Order Date:%s", orderDetails.OrderDate))
	pdf.Ln(15)

	// Add user details
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 6, "User Details:")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	userDetails := fmt.Sprintf("%s %s\n%s\n%s", orderDetails.FName, orderDetails.LName, orderDetails.Email, orderDetails.Phone)
	pdf.MultiCell(0, 5, userDetails, "", "", false)
	pdf.Ln(-1)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 6, "Delivery Address:")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	addressDetails := fmt.Sprintf("%s\n%s,%s,%s\n%s\n%s\n%s", orderDetails.Line1, orderDetails.Street, orderDetails.City, orderDetails.State, orderDetails.PostalCode, orderDetails.Country, orderDetails.AlternatePhone)
	pdf.MultiCell(0, 5, addressDetails, "", "", false)

	pdf.Ln(-1)

	// Table header
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "ORDERED ITEMS:")
	pdf.Ln(-1)

	// Table headers
	headers := []string{"ITEM", "QUANTITY", "RESTAURANT", "MRP", "P-DISCOUNT", "C-DISCOUNT", "SALE-PRICE", "AMOUNT"}

	pdf.SetFont("Arial", "B", 8)

	// Print table headers for
	for _, header := range headers {
		pdf.CellFormat(23, 6, header, "1", 0, "", false, 0, "")
	}

	pdf.Ln(-1)

	tableData := [][]string{}

	// Print table data
	pdf.SetFont("Arial", "", 6)

	// Populate tableData dynamically from the fetched data
	for _, item := range *orderedItemsDetails {
		applP := item.MRP - ((item.MRP * float64(item.PromotionDiscount)) / 100)
		applC := applP - ((applP * item.DiscountPercentage) / 100)
		totalAmount = float64(item.OrderQuantity) * applC
		FinalAmount += totalAmount
		row := []string{
			item.Name,
			fmt.Sprintf("%d", item.OrderQuantity),
			item.Restaurant_name,
			fmt.Sprintf("%.2f", item.MRP),
			fmt.Sprintf("%d%%", item.PromotionDiscount),
			fmt.Sprintf("%.0f%%", item.DiscountPercentage),
			fmt.Sprintf("%.2f", applC),
			fmt.Sprintf("%.2f", totalAmount),
		}
		tableData = append(tableData, row)
	}

	for _, line := range tableData {
		for _, cell := range line {
			pdf.CellFormat(23, 5, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	pdf.SetFont("Arial", "B", 8)
	// Add coupon details
	if orderDetails.CouponCode == "" {
		orderDetails.CouponCode = "Nil"
	}
	pdf.Ln(-1)
	pdf.SetX(115)
	pdf.CellFormat(0, 5, "-----------------------------------------------------------------------------------", "", 1, "L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(0, 5, fmt.Sprintf("Final Amount: %.2f", FinalAmount), "", 1, "L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(0, 5, "Coupon Applied: "+orderDetails.CouponCode, "", 1, "L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(0, 5, fmt.Sprintf("Discount: %.2f%%", orderDetails.DiscountPercentage), "", 1, "L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(0, 5, fmt.Sprintf("Final amount after applying coupon and round-off: %.2f", math.Ceil(FinalAmount-((FinalAmount*orderDetails.DiscountPercentage)/100))), "", 1, "L", false, 0, "")

	// Add "Thank you" message
	pdf.SetY(220)
	pdf.SetX(70)
	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(0, 10, "Thank you for shopping with us!")

	var pdfbuffer bytes.Buffer
	errGener := pdf.Output(&pdfbuffer)
	if errGener != nil {
		return &pdfLink, errGener
	}

	sess, errInit := aws.AWSSessionInitializer()
	if errInit != nil {
		fmt.Println(errInit)
		return &pdfLink, errInit
	}
	generatedURL, errS3 := aws.AWSS3ByteBufferUploader(pdfbuffer.Bytes(), sess, &BucketFolder)
	if errS3 != nil {
		fmt.Println("Error uploading to S3:", err)
		return &pdfLink, errS3
	}

	return generatedURL, nil
}

func (r *OrderUseCase) GenerateSalesReportXlsx(restId *string) (*string, error) {
	var dummyString *string
	BucketFolder := "vegn-ecommerce-tempfiles/sales-reports/"

	orderData, err := r.OrderRepo.OrdersForSalesReportByRestId(restId)
	if err != nil {
		return dummyString, err
	}

	f := excelize.NewFile()

	sheetName := "SalesReport"
	intret := f.NewSheet(sheetName)

	// Set column headers
	headers := []string{"OrderID", "DishID", "DishName", "Quantity", "MRP", "PromotionDiscount", "CategoryDiscount", "SalePrice", "PaymentMethod", "OrderDate", "DeliverDate"}
	for colIndex, header := range headers {
		cell := excelize.ToAlphaString(colIndex+1) + "1"
		f.SetCellValue(sheetName, cell, header)
	}

	// Populate the sheet with data
	for rowIndex, record := range *orderData {
		colIndex := 1
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.OrderID)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.DishID)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.Name)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.OrderQuantity)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.MRP)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.PromotionDiscount)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.DiscountPercentage)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.DishPrice)
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.PaymentMethod)

		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.OrderDate.Format("2006-01-02 15:04:05"))
		colIndex++
		f.SetCellValue(sheetName, excelize.ToAlphaString(colIndex)+fmt.Sprint(rowIndex+2), record.DeliverDate.Format("2006-01-02 15:04:05"))
	}

	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(intret)

	// Save the Excel file
	var xlsxbuffer bytes.Buffer
	errSave := f.Write(&xlsxbuffer)
	if errSave != nil {
		return dummyString, errSave
	}

	sess, errInit := aws.AWSSessionInitializer()
	if errInit != nil {
		fmt.Println(errInit)
		return dummyString, errInit
	}

	generatedURL, errS3 := aws.AWSS3XLSXUploader(xlsxbuffer.Bytes(), sess, &BucketFolder)
	if errS3 != nil {
		fmt.Println("Error uploading to S3:", err)
		return dummyString, errS3
	}
	return generatedURL, nil
}

func (r *OrderUseCase) GetSalesReportForCustomDays(restaurantId *string, customDays *string) (*responsemodels.SalesReportData, error) {

	dataStruct, errChk := r.OrderRepo.GetDataSalesReportForCustomDays(restaurantId, customDays)
	if errChk != nil {
		return dataStruct, errChk
	}

	return dataStruct, nil

}

func (r *OrderUseCase) GetSalesreporYYMMDD(restaurantId *string, yymmdd *requestmodels.SalesReportYYMMDD) (*responsemodels.SalesReportData, error) {

	validate := validator.New(validator.WithRequiredStructEnabled())
	errValidate := validate.Struct(yymmdd)
	if errValidate != nil {
		return nil, errValidate
	}

	dataStruct, errChk := r.OrderRepo.GetDataSalesReportYYMMDD(restaurantId, yymmdd)
	if errChk != nil {
		return dataStruct, errChk
	}

	return dataStruct, nil
}
