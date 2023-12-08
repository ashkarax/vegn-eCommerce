package usecase

import (
	"errors"
	"fmt"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/ashkarax/vegn-eCommerce/pkg/razorpay"
)

type PaymentUsecase struct {
	OrderRepo interfaceRepository.IOrderRepo
	CartRepo  interfaceRepository.ICartRepository
	RazorKeys config.RazorPay
}

func NewPaymentUsecase(orderRepo interfaceRepository.IOrderRepo, cartRepo interfaceRepository.ICartRepository, razorKeys *config.RazorPay) interfaceUseCase.IPaymentUseCase {
	return &PaymentUsecase{OrderRepo: orderRepo, CartRepo: cartRepo, RazorKeys: *razorKeys}
}

func (r *PaymentUsecase) GetOrderDetails(userId *string, orderId *string) (*responsemodels.RazorpayResponse, error) {
	var razorResp responsemodels.RazorpayResponse
	var totalPrice float64

	razorResp.KeyId = r.RazorKeys.KeyId

	orderItemsDetails, errfetch := r.OrderRepo.GetOrderDetailsByOrderId(userId, orderId)
	if errfetch != nil {
		fmt.Println("-----------", errfetch)
	}

	for _, item := range *orderItemsDetails {
		totalPrice += float64(item.OrderQuantity) * item.DishPrice
		fullName := item.FName + " " + item.LName

		razorResp.UserName = fullName
		razorResp.RazorPayId = item.RazorPayId
		razorResp.Email = item.Email
		razorResp.Phone = item.Phone

	}
	razorResp.TotalAmount = totalPrice * 100

	return &razorResp, nil

}

func (r *PaymentUsecase) OnlinePaymentVerification(razorResp *requestmodels.RazorWebOut) (*[]responsemodels.OrderDetailsResponse, error) {
	var dumbSlice []responsemodels.OrderDetailsResponse
	razorResp.RSecrect = r.RazorKeys.SecrectKey

	if razorpay.VerifyPayment(razorResp) {
		orderDetails, err := r.OrderRepo.UpdateStatusToSuccess(&razorResp.UserId, &razorResp.OrderId, &razorResp.ROrderId)
		if err != nil {
			return orderDetails, err
		}

		return orderDetails, nil
	}

	return &dumbSlice, errors.New("payment verification failed")

}
