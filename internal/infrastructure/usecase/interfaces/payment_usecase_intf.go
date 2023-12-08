package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IPaymentUseCase interface {
	GetOrderDetails(*string, *string) (*responsemodels.RazorpayResponse, error)
	OnlinePaymentVerification(razorResp *requestmodels.RazorWebOut) (*[]responsemodels.OrderDetailsResponse, error)
}
