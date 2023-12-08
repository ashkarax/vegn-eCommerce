package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IOrderRepo interface {
	PlaceNewOrder(*requestmodels.OrderDetails) (*string, error)
	AddDishesToOrderedItems(*string, *responsemodels.CartItemInfo, *requestmodels.OrderDetails) error

	GetAllOrdersByUser(userId *string) (*[]responsemodels.OrderDetailsResponse, error)

	OrdersForRestaurantById(*string) (*[]responsemodels.OrderDetailsResponse, error)

	GetOrderDetailsByOrderId(*string, *string) (*[]responsemodels.RazorpayResponse, error)

	UpdateStatusToSuccess(*string, *string, *string) (*[]responsemodels.OrderDetailsResponse, error)

	ReturnOrderStats(*requestmodels.CanOrRetReq) (*responsemodels.CanOrRetResp, error)
	ChangeOrderStatus(*string, *string) error
	UpdatePaymentStatus(*string, *string) error

	RestReturnOrderStatus(*string, *string) (*responsemodels.CanOrRetResp, error)

	UpdateDeliveryDate(*string)
}
