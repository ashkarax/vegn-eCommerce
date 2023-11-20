package handlers

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	sellerUsecase interfaceUseCase.IsellerUseCase
}

func NewSellerHandler(sellerUsecase interfaceUseCase.IsellerUseCase) *SellerHandler {
	return &SellerHandler{sellerUsecase: sellerUsecase}
}

func (u *SellerHandler) SellerSignUp(c *gin.Context) {
	var sellerSignUpData requestmodels.SellerSignUpReq
	
	if err := c.BindJSON(&sellerSignUpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resSignup, err := u.sellerUsecase.SellerSignUp(&sellerSignUpData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Signup Failed", resSignup, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully signup", resSignup, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *SellerHandler) SellerLogin(c *gin.Context) {
	var sellerLoginData requestmodels.SellerLoginReq

	if err := c.BindJSON(&sellerLoginData); err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	sellerLogin, err := u.sellerUsecase.SellerLogin(&sellerLoginData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "Login Failed", sellerLogin, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully signup", sellerLogin, nil)
	c.JSON(http.StatusOK, finalReslt)
}
