package handlers

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	AddressUseCase interfaceUseCase.IAddressUseCase
}

func NewAddressHandler(addressUseCase interfaceUseCase.IAddressUseCase) *AddressHandler {
	return &AddressHandler{AddressUseCase: addressUseCase}
}

func (u *AddressHandler) AddNewAddress(c *gin.Context) {
	var newAddressData requestmodels.AddressReq

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	newAddressData.UserID = userIdString

	if err := c.BindJSON(&newAddressData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resAddressData, err := u.AddressUseCase.AddNewAddress(&newAddressData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add new Address", resAddressData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Address added succesfully", resAddressData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *AddressHandler) EditAddress(c *gin.Context) {
	var editAddressData requestmodels.AddressReq

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	editAddressData.UserID = userIdString

	dishid := c.Param("addressId")
	editAddressData.AddressId = dishid

	if err := c.BindJSON(&editAddressData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resEditAddressData, err := u.AddressUseCase.EditAddress(&editAddressData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't update Address", resEditAddressData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Address updated succesfully", resEditAddressData, nil)
	c.JSON(http.StatusOK, response)

}

func (u *AddressHandler) GetAllAddress(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	resData, err := u.AddressUseCase.GetAllAddresses(&userIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "error finding addresses", resData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Addresses fetched succesfully", resData, nil)
	c.JSON(http.StatusOK, response)
}
