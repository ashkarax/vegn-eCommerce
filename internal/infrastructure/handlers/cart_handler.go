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

//	@Summary		AddToCart
//	@Description	Adds a dish to the user's cart.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Security		UserAuthTokenAuth
//	@Security		UserRefTokenAuth
//	@Param			dishid	path		string	true	"The ID of the dish to add."
//	@Success		200		{object}	responsemodels.Response
//	@Failure		400		{object}	responsemodels.Response
//	@Router			/cart/addtocart/{dishid} [post]
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

//	@Summary		GetCartDetailsOfUser
//	@Description	Retrieves the cart details for the current user.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Security		UserAuthTokenAuth
//	@Security		UserRefTokenAuth
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/cart [get]
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

//	@Summary		DecrementorRemoveFromCart
//	@Description	Removes a dish from the user's cart.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Security		UserAuthTokenAuth
//	@Security		UserRefTokenAuth
//	@Param			dishid	path		string	true	"The ID of the dish to remove."
//	@Success		200		{object}	responsemodels.Response
//	@Failure		400		{object}	responsemodels.Response
//	@Router			/cart/{dishid} [delete]
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
