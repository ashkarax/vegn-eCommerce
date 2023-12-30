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
