package handlers

import (
	"fmt"
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	UseCase interfaceUseCase.ICategoryUseCase
}

func NewCategoryHandler(useCase interfaceUseCase.ICategoryUseCase) *CategoryHandler {
	return &CategoryHandler{UseCase: useCase}
}

//	@Summary		NewCategory
//	@Description	Creates a new category.
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Security		AdminRefTokenAuth  
//	@Param			category	body		requestmodels.CategoryReq	true	"New category data."
//	@Success		200			{object}	responsemodels.CategoryRes
//	@Failure		400			{object}	responsemodels.Response
//	@Router			/admin/category [post]
func (u *CategoryHandler) NewCategory(c *gin.Context) {
	var category requestmodels.CategoryReq
	var categoryRes responsemodels.CategoryRes

	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryId, err := u.UseCase.AddNewCategory(&category)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add category", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	categoryRes.CategoryID = *categoryId

	response := responsemodels.Responses(http.StatusOK, "category added succesfully", categoryRes, nil)
	c.JSON(http.StatusOK, response)
}

//	@Summary		FetchAllCategory
//	@Description	Retrieves all available categories.
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Security		AdminRefTokenAuth  
//	@Success		200	{array}		responsemodels.CategoryRes
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/admin/category [get]
func (u *CategoryHandler) FetchAllCategory(c *gin.Context) {
	categorySlice, err := u.UseCase.GetAllCategories()
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed fetching categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "categories fetched succesfully", categorySlice, nil)
	c.JSON(http.StatusOK, response)
}

//	@Summary		DeleteCategory
//	@Description	Deletes a category.
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Security		AdminRefTokenAuth
//	@Param			categoryid	path		string	true	"The ID of the category to delete."
//	@Success		200			{object}	responsemodels.Response
//	@Failure		400			{object}	responsemodels.Response
//	@Router			/admin/category/{categoryid} [delete]
func (u *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryid := c.Param("categoryid")
	newStatus := "deleted"
	err := u.UseCase.ChangeCategoryStatus(&categoryid, &newStatus)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete category", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := responsemodels.Responses(http.StatusOK, "category deleted succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

//	@Summary		UpdateCategory
//	@Description	Updates an existing category.
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Security		AdminRefTokenAuth
//	@Param			categoryid	path		string						true	"The ID of the category to update."
//	@Param			category	body		requestmodels.CategoryReq	true	"Updated category data."
//	@Success		200			{object}	responsemodels.CategoryRes
//	@Failure		400			{object}	responsemodels.Response
//	@Router			/admin/category/{categoryid} [patch]
func (u *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryid := c.Param("categoryid")
	var category requestmodels.CategoryReq
	var categoryRes responsemodels.CategoryRes

	
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.CategoryId = categoryid
	categoryId, err := u.UseCase.UpdateCategory(&category)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't update category", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	categoryRes.CategoryID = *categoryId

	response := responsemodels.Responses(http.StatusOK, "category updated succesfully", categoryRes, nil)
	c.JSON(http.StatusOK, response)
}

//	@Summary		Fetch active categories
//	@Description	Fetch all active categories
//	@Tags			RestaurantCategoryManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/restaurant/category [get]
func (u *CategoryHandler) FetchActiveCategories(c *gin.Context) {
	categorySlice, err := u.UseCase.FetchActiveCategories()
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed fetching categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "categories fetched succesfully", categorySlice, nil)
	c.JSON(http.StatusOK, response)
}

//	@Summary		Create a new category offer
//	@Description	Create a new category offer for the restaurant specified in the request context
//	@Tags			RestaurantCategoryOfferManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Param			categoryOfferData	body		requestmodels.CategoryOfferReq	true	"Category offer information"
//	@Success		200					{object}	responsemodels.Response
//	@Failure		400					{object}	responsemodels.Response
//	@Router			/restaurant/category/offer [post]
func (u *CategoryHandler) CreateCategoryOffer(c *gin.Context) {

	var categoryOffer requestmodels.CategoryOfferReq
	var categoryOfferRes responsemodels.CategoryOfferRes

	

	if err := c.BindJSON(&categoryOffer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryOffer.RestaurantID = c.MustGet("RestaurantId").(string)

	categoryOfferId, err := u.UseCase.AddNewCategoryOffer(&categoryOffer)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	categoryOfferRes.CategoryOfferID = *categoryOfferId

	response := responsemodels.Responses(http.StatusOK, "category offer added succesfully", categoryOfferRes, nil)
	c.JSON(http.StatusOK, response)

}

//	@Summary		Fetch all category offers
//	@Description	Fetch all category offers for the restaurant specified in the request context
//	@Tags			RestaurantCategoryOfferManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/restaurant/category/offer [get]
func (u *CategoryHandler) GetAllCategoryOffer(c *gin.Context) {

	restaurantId := c.MustGet("RestaurantId").(string)

	categoryOfferSlice, err := u.UseCase.GetAllCategoryOffers(&restaurantId)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed fetching category offers", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "category offers fetched succesfully", categoryOfferSlice, nil)
	c.JSON(http.StatusOK, response)

}

//	@Summary		Edit a category offer
//	@Description	Edit a category offer for the restaurant specified in the request context
//	@Tags			RestaurantCategoryOfferManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Param			categoryofferid		path		string							true	"Category Offer ID to edit."
//	@Param			categoryOfferData	body		requestmodels.EditCategoryOffer	true	"Updated category offer information"
//	@Success		200					{object}	responsemodels.Response
//	@Failure		400					{object}	responsemodels.Response
//	@Router			/restaurant/category/offer/{categoryofferid} [patch]
func (u *CategoryHandler) EditCategoryOffer(c *gin.Context) {

	var categoryOffer requestmodels.EditCategoryOffer

	

	if err := c.BindJSON(&categoryOffer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryOffer.CategoryOfferID = c.Param("categoryofferid")
	categoryOffer.RestaurantID = c.MustGet("RestaurantId").(string)

	updatedCategoryOffer, err := u.UseCase.EditCategoryOffer(&categoryOffer)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't edit category", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "category edited succesfully", updatedCategoryOffer, nil)
	c.JSON(http.StatusOK, response)

}

//	@Summary		Change the status of a category offer
//	@Description	Change the status of a category offer for the restaurant specified in the request context
//	@Tags			RestaurantCategoryOfferManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Param			categoryOfferStatusData	body		requestmodels.CategoryOfferStatus	true	"Category offer status information"
//	@Success		200						{object}	responsemodels.Response
//	@Failure		400						{object}	responsemodels.Response
//	@Router			/restaurant/category/offer/ [patch]
func (u *CategoryHandler) ChangeCategoryOfferStatus(c *gin.Context) {
	var status requestmodels.CategoryOfferStatus
	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.UseCase.ChangeCategoryOfferStatus(&status)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't change category status", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resp := fmt.Sprintf("category offer status with id %s changed to %s succesfully", status.CategoryOfferId, status.Status)
	response := responsemodels.Responses(http.StatusOK, resp, nil, nil)
	c.JSON(http.StatusOK, response)

}
