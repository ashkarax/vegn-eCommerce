package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderHandler struct {
	OrderUseCase interfaceUseCase.IOrderUseCase
}

func NewOrderHandler(orderUseCase interfaceUseCase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{OrderUseCase: orderUseCase}
}

// @Summary PlaceNewOrder
// @Description Places a new order for the current user.
// @Tags Orders
// @Accept json
// @Produce json
// @Security UserAuthTokenAuth
// @Security UserRefTokenAuth
// @Param orderDetails body requestmodels.OrderDetails true "Order details."
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /order [post]
func (u *OrderHandler) PlaceNewOrder(c *gin.Context) {
	var orderDetails requestmodels.OrderDetails

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	orderDetails.UserID = userIdString
	if err := c.BindJSON(&orderDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := u.OrderUseCase.PlaceNewOrder(&orderDetails)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to place your order", resp, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "order placed succesfully", resp, nil)
	c.JSON(http.StatusOK, response)
}

// @Summary GetAllOrders
// @Description Retrieves all orders for the current user.
// @Tags Orders
// @Accept json
// @Produce json
// @Security UserAuthTokenAuth
// @Security UserRefTokenAuth
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /order [get]
func (u *OrderHandler) GetAllOrders(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	orderMap, err := u.OrderUseCase.GetAllOrders(&userIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to find user's orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "orders fetched succesfully", orderMap, nil)
	c.JSON(http.StatusOK, response)
}

// @Summary FetchAllOrdersForRestaurant
// @Description Retrieves all orders for a specific restaurant.
// @Tags RestaurantOrderManagement
// @Accept json
// @Produce json
// @Security RestaurantAuthTokenAuth
// @Security RestaurantRefTokenAuth
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /restaurant/orders [get]
func (u *OrderHandler) FetchAllOrdersForRestaurant(c *gin.Context) {
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdString, _ := restaurantId.(string)

	orderMap, err := u.OrderUseCase.AllOrdersForARestaurant(&restaurantIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to find restaurant's orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "orders fetched succesfully", orderMap, nil)
	c.JSON(http.StatusOK, response)
}

// func (u *OrderHandler) FetchOrderWithId(c *gin.Context) {
// 	restaurantId, _ := c.Get("RestaurantId")
// 	restaurantIdString, _ := restaurantId.(string)

// 	orderid := c.Param("orderid")

// 	orderDetails,err := u.OrderUseCase.FetchOrderForRestaurant(&restaurantIdString,&orderid)
// 	if err != nil {
// 		response := responsemodels.Responses(http.StatusBadRequest, "failed to find restaurant's orders", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := responsemodels.Responses(http.StatusOK, "orders fetched succesfully", orderDetails, nil)
// 	c.JSON(http.StatusOK, response)

// }
func (u *OrderHandler) CancelOrder(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	ordereditemsid := c.Param("ordereditemsid")

	var orderDetails requestmodels.CanOrRetReq
	orderDetails.OrderedItemsID = ordereditemsid
	orderDetails.UserID = userIdString

	err := u.OrderUseCase.CancelOrderById(&orderDetails)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to find order", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	endres := fmt.Sprintf("order having ordereditemsid=%s cancelled succesfully", ordereditemsid)
	response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
	c.JSON(http.StatusOK, response)

}

func (u *OrderHandler) ReturnOrder(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	ordereditemsid := c.Param("ordereditemsid")

	var orderDetails requestmodels.CanOrRetReq
	orderDetails.OrderedItemsID = ordereditemsid
	orderDetails.UserID = userIdString

	err := u.OrderUseCase.ReturnOrderById(&orderDetails)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to find order", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	endres := fmt.Sprintf("order having ordereditemsid=%s cancelled succesfully", ordereditemsid)
	response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
	c.JSON(http.StatusOK, response)

}

// @Summary ChangeStatus
// @Description Updates the status of an order.
// @Tags RestaurantOrderManagement
// @Accept json
// @Produce json
// @Security RestaurantAuthTokenAuth
// @Security RestaurantRefTokenAuth
// @Param changeStat body requestmodels.ChangeStatus true "Details of the status change."
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /restaurant/orders/status [patch]
func (u *OrderHandler) ChangeStatus(c *gin.Context) {
	var changeStat requestmodels.ChangeStatus

	restaurantId, _ := c.Get("RestaurantId")
	changeStat.RestaurantId = restaurantId.(string)

	if err := c.BindJSON(&changeStat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	errValidate := validate.Struct(changeStat)
	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidate.Error()})
		return
	}

	if changeStat.Status == "preparing" {
		err := u.OrderUseCase.ChangeStatusToPreparing(&changeStat.RestaurantId, &changeStat.OrderedItemsID)
		if err != nil {
			response := responsemodels.Responses(http.StatusBadRequest, "failed to change order status", nil, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
		endres := fmt.Sprintf("order having ordereditemsid=%s status changed to preparing succesfully", changeStat.OrderedItemsID)
		response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
		c.JSON(http.StatusOK, response)
		return
	}

	if changeStat.Status == "outfordelivery" {
		err := u.OrderUseCase.ChangeStatusToOutForDelivery(&changeStat.RestaurantId, &changeStat.OrderedItemsID)
		if err != nil {
			response := responsemodels.Responses(http.StatusBadRequest, "failed to change order status", nil, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
		endres := fmt.Sprintf("order having ordereditemsid=%s status changed to out-for-delivery succesfully", changeStat.OrderedItemsID)
		response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
		c.JSON(http.StatusOK, response)
		return

	}

	if changeStat.Status == "deliver" {
		err := u.OrderUseCase.ChangeStatusToDelivered(&changeStat.RestaurantId, &changeStat.OrderedItemsID)
		if err != nil {
			response := responsemodels.Responses(http.StatusBadRequest, "failed to change order status", nil, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
		endres := fmt.Sprintf("order having ordereditemsid=%s status changed to delivered succesfully", changeStat.OrderedItemsID)
		response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := responsemodels.Responses(http.StatusBadRequest, "enter a valid status", nil, nil)
	c.JSON(http.StatusBadRequest, response)

}

// @Summary GenerateInvoice
// @Description Generates an invoice for a specific order.
// @Tags Orders
// @Accept json
// @Produce json
// @Security UserAuthTokenAuth
// @Security UserRefTokenAuth
// @Param orderid path string true "The ID of the order to generate an invoice for."
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /order/invoice/{orderid} [get]
func (u *OrderHandler) GenerateInvoice(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	orderid := c.Param("orderid")

	pdfURL, err := u.OrderUseCase.GenerateInvoice(&userIdString, &orderid)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to find order", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	endres := fmt.Sprintf("order having ordereditemsid=%s's invoice generated succesfully", orderid)
	response := responsemodels.Responses(http.StatusOK, endres, pdfURL, nil)
	c.JSON(http.StatusOK, response)

}


// @Summary GenerateSalesReportXlsx
// @Description Generates a sales report in XLSX format.
// @Tags RestaurantReports
// @Accept json
// @Produce json
// @Security RestaurantAuthTokenAuth
// @Security RestaurantRefTokenAuth
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /restaurant/salesreport/xlsx [get]
func (u *OrderHandler) GenerateSalesReportXlsx(c *gin.Context) {
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdtring, _ := restaurantId.(string)

	xlsxURL, err := u.OrderUseCase.GenerateSalesReportXlsx(&restaurantIdtring)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to generate Sales Report", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "SalesReport generated succesfully", xlsxURL, nil)
	c.JSON(http.StatusOK, response)

}

// @Summary GetSalesreporCustomDays
// @Description Generates a sales report for a custom number of days.
// @Tags RestaurantReports
// @Accept json
// @Produce json
// @Security RestaurantAuthTokenAuth
// @Security RestaurantRefTokenAuth
// @Param customdays path int true "Number of recent days for the report."
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /restaurant/salesreport/{customdays} [get]
func (u *OrderHandler) GetSalesreporCustomDays(c *gin.Context) {
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdtring, _ := restaurantId.(string)

	customDays := c.Param("customdays")
	_, errConv := strconv.Atoi(customDays)

	if customDays == "" || errConv != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to generate Sales Report", errConv, errors.New("customDays must be a valid number"))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	returnData, err := u.OrderUseCase.GetSalesReportForCustomDays(&restaurantIdtring, &customDays)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to generate Sales Report", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "SalesReport generated succesfully", returnData, nil)
	c.JSON(http.StatusOK, response)

}

// @Summary GetSalesreporYYMMDD
// @Description Generates a sales report for a specific date range or a specific month or a specific year.
// @Tags RestaurantReports
// @Accept json
// @Produce json
// @Security RestaurantAuthTokenAuth
// @Security RestaurantRefTokenAuth
// @Param yymmdd body requestmodels.SalesReportYYMMDD true "YYMMDD for the report ."
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /restaurant/salesreport [post]
func (u *OrderHandler) GetSalesreporYYMMDD(c *gin.Context) {
	var yymmdd requestmodels.SalesReportYYMMDD
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdtring, _ := restaurantId.(string)

	if err := c.BindJSON(&yymmdd); err != nil {
		if err.Error() == "EOF"{
			c.JSON(http.StatusBadRequest, gin.H{"error": "no json input found in your request"})
		return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	returnData, err := u.OrderUseCase.GetSalesreporYYMMDD(&restaurantIdtring, &yymmdd)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to generate Sales Report", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "SalesReport generated succesfully", returnData, nil)
	c.JSON(http.StatusOK, response)

}
