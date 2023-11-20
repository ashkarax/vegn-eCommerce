package handlers

import (
	"fmt"
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUseCase interfaceUseCase.IAdminUseCase
}

func NewAdminHandler(useCase interfaceUseCase.IAdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUseCase: useCase}
}

func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var loginCredential requestmodels.AdminLoginData

	bindErr := c.BindJSON(&loginCredential)
	if bindErr != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "json is wrong can't bind", nil, bindErr.Error())
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	result, validateErr := u.AdminUseCase.AdminLogin(&loginCredential)
	if validateErr != nil {
		fmt.Println("", validateErr)
		finalReslt := responsemodels.Responses(http.StatusUnauthorized, "Not adnim with this email", nil, validateErr.Error())
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully login", result, nil)
	c.JSON(http.StatusOK, finalReslt)
}
