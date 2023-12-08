package handlers

import (
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

func (u *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryid := c.Param("categoryid")
	var category requestmodels.CategoryReq
	var categoryRes responsemodels.CategoryRes

	category.CategoryId = categoryid
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
