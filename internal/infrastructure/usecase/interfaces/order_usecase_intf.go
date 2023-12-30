package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IOrderUseCase interface {
	PlaceNewOrder(*requestmodels.OrderDetails) (*responsemodels.OrderDetailsRes, error)
	GetAllOrders(*string) (*[]responsemodels.OrderDetailsResponse, error)

	AllOrdersForARestaurant(*string) (*[]responsemodels.OrderResponseX, error)

	CancelOrderById(*requestmodels.CanOrRetReq) error
	ReturnOrderById(*requestmodels.CanOrRetReq) error

	ChangeStatusToPreparing(*string, *string) error
	ChangeStatusToOutForDelivery(*string, *string) error
	ChangeStatusToDelivered(*string, *string) error

	GenerateInvoice(*string, *string) (*string, error)

	GenerateSalesReportXlsx(*string) (*string, error)

	GetSalesReportForCustomDays(*string, *string) (*responsemodels.SalesReportData, error)

	GetSalesreporYYMMDD(*string, *requestmodels.SalesReportYYMMDD) (*responsemodels.SalesReportData, error)
}
