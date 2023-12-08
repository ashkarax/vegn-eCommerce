package handlers

import (
	"fmt"
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
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

func (u *OrderHandler) ChangeStatusToPreparing(c *gin.Context) {
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdString, _ := restaurantId.(string)
	ordereditemsid := c.Param("ordereditemsid")
	err := u.OrderUseCase.ChangeStatusToPreparing(&restaurantIdString, &ordereditemsid)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to change order status", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	endres := fmt.Sprintf("order having ordereditemsid=%s status changed to preparing succesfully", ordereditemsid)
	response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *OrderHandler) ChangeStatusToOutForDelivery(c *gin.Context) {
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdString, _ := restaurantId.(string)
	ordereditemsid := c.Param("ordereditemsid")
	err := u.OrderUseCase.ChangeStatusToOutForDelivery(&restaurantIdString, &ordereditemsid)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to change order status", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	endres := fmt.Sprintf("order having ordereditemsid=%s status changed to out-for-delivery succesfully", ordereditemsid)
	response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *OrderHandler) ChangeStatusToDelivered(c *gin.Context) {
	restaurantId, _ := c.Get("RestaurantId")
	restaurantIdString, _ := restaurantId.(string)
	ordereditemsid := c.Param("ordereditemsid")
	err := u.OrderUseCase.ChangeStatusToDelivered(&restaurantIdString, &ordereditemsid)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to change order status", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	endres := fmt.Sprintf("order having ordereditemsid=%s status changed to delivered succesfully", ordereditemsid)
	response := responsemodels.Responses(http.StatusOK, endres, nil, nil)
	c.JSON(http.StatusOK, response)
}

// ordermanagement.PATCH("/:orderid/cancel",order.RestaurantCancelOrder)
