package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IOrderUseCase interface {
	PlaceNewOrder(*requestmodels.OrderDetails) (*responsemodels.OrderDetailsRes, error)
	GetAllOrders(*string) (*[]responsemodels.OrderDetailsResponse, error)

	AllOrdersForARestaurant(*string) (*[]responsemodels.OrderDetailsResponse, error)

	CancelOrderById(*requestmodels.CanOrRetReq) error
	ReturnOrderById(*requestmodels.CanOrRetReq) error

	ChangeStatusToPreparing(*string, *string) error
	ChangeStatusToOutForDelivery(*string, *string) error
	ChangeStatusToDelivered(*string, *string) error


	
}
