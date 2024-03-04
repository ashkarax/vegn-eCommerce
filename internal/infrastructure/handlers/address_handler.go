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

// @Summary		AddNewAddress.
// @Description	Adds a new address for the current user,updated today
// @Tags			Address
// @Accept			json
// @Produce		json
// @Security		UserAuthTokenAuth
// @Security		UserRefTokenAuth
// @Param			newAddressData	body		requestmodels.AddressReq	true	"New address data."
// @Success		200				{object}	responsemodels.Response
// @Failure		400				{object}	responsemodels.Response
// @Router			/address [post]
func (u *AddressHandler) AddNewAddress(c *gin.Context) {
	var newAddressData requestmodels.AddressReq

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	if err := c.BindJSON(&newAddressData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAddressData.UserID = userIdString

	resAddressData, err := u.AddressUseCase.AddNewAddress(&newAddressData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add new Address", resAddressData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Address added succesfully", resAddressData, nil)
	c.JSON(http.StatusOK, response)
}

// @Summary		EditAddress
// @Description	Edits an existing address for the current user.
// @Tags			Address
// @Accept			json
// @Produce		json
// @Security		UserAuthTokenAuth
// @Security		UserRefTokenAuth
// @Param			addressId		path		string						true	"The ID of the address to edit."
// @Param			editAddressData	body		requestmodels.AddressReq	true	"Address edit data."
// @Success		200				{object}	responsemodels.Response
// @Failure		400				{object}	responsemodels.Response
// @Router			/address/{addressId} [patch]
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

// @Summary		GetAllAddress
// @Description	Retrieves all addresses for the current user.
// @Tags			Address
// @Accept			json
// @Produce		json
// @Security		UserAuthTokenAuth
// @Security		UserRefTokenAuth
// @Success		200	{object}	responsemodels.Response
// @Failure		400	{object}	responsemodels.Response
// @Router			/address [get]
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
