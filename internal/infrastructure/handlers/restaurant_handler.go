package handlers

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	RestaurantUsecase interfaceUseCase.IrestaurantUseCase
}

func NewRestaurantHandler(RestaurantUsecase interfaceUseCase.IrestaurantUseCase) *RestaurantHandler {
	return &RestaurantHandler{RestaurantUsecase: RestaurantUsecase}
}

func (u *RestaurantHandler) RestaurantSignUp(c *gin.Context) {
	var restaurantSignUpData requestmodels.RestaurantSignUpReq
	
	if err := c.BindJSON(&restaurantSignUpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resSignup, err := u.RestaurantUsecase.RestaurantSignUp(&restaurantSignUpData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Signup Failed", resSignup, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully signup", resSignup, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *RestaurantHandler) RestaurantLogin(c *gin.Context) {
	var RestaurantLoginData requestmodels.RestaurantLoginReq

	if err := c.BindJSON(&RestaurantLoginData); err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	RestaurantLogin, err := u.RestaurantUsecase.RestaurantLogin(&RestaurantLoginData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "Login Failed", RestaurantLogin, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully signup", RestaurantLogin, nil)
	c.JSON(http.StatusOK, finalReslt)
}
