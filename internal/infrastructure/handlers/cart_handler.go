package handlers

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	UseCase interfaceUseCase.ICartUseCase
}

func NewCartHandler(useCase interfaceUseCase.ICartUseCase) *CartHandler {
	return &CartHandler{UseCase: useCase}
}

func (u *CartHandler) AddToCart(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	dishId := c.Param("dishid")

	err := u.UseCase.AddToCart(&userIdString, &dishId)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "error adding to cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Addded to cart succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *CartHandler) GetCartDetailsOfUser(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	cartDetails, err := u.UseCase.GetCartDetails(&userIdString)

	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "error fetching cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "cart details fetched succesfully", cartDetails, nil)
	c.JSON(http.StatusOK, response)

}

func (u *CartHandler) DecrementorRemoveFromCart(c *gin.Context) {
	dishid := c.Param("dishid")
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	err := u.UseCase.DeleteFromCart(&dishid, &userIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "error removing dish from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "succesfully removed item from cart", nil, nil)
	c.JSON(http.StatusOK, response)
}
