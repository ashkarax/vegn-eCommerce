package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type DishHandler struct {
	DishUseCase interfaceUseCase.IDishUseCase
}

func NewDishHandler(dishUseCase interfaceUseCase.IDishUseCase) *DishHandler {
	return &DishHandler{DishUseCase: dishUseCase}
}

func (u *DishHandler) NewDish(c *gin.Context) {
	var newDishData requestmodels.DishReq

	RestaurantId, _ := c.Get("RestaurantId")
	RestaurantIdString, _ := RestaurantId.(string)
	newDishData.RestaurantId = RestaurantIdString

	if err := c.ShouldBind(&newDishData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("------------", newDishData)
	resDishData, err := u.DishUseCase.NewDish(&newDishData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add new Dish", resDishData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "dish added succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *DishHandler) FetchAllDishesForRestaurant(c *gin.Context) {
	RestaurantId, _ := c.Get("RestaurantId")
	RestaurantIdString, _ := RestaurantId.(string)
	num, _ := strconv.Atoi(RestaurantIdString)

	dishSlice, err := u.DishUseCase.FetchAllDishesForRestaurant(&num)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch dishes", dishSlice, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := responsemodels.Responses(http.StatusOK, "dishes fetched succesfully", dishSlice, nil)
	c.JSON(http.StatusOK, response)
}
func (u *DishHandler) FetchDishWithId(c *gin.Context) {
	dishid := c.Param("dishid")
	num, _ := strconv.Atoi(dishid)

	dishData, err := u.DishUseCase.DishById(&num)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch dish", dishData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := responsemodels.Responses(http.StatusOK, "dish fetched succesfully", dishData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *DishHandler) UpdateDishDetails(c *gin.Context) {
	dishid := c.Param("dishid")
	num, _ := strconv.Atoi(dishid)

	var dishData requestmodels.DishUpdateReq

	RestaurantId, _ := c.Get("RestaurantId")
	RestaurantIdString, _ := RestaurantId.(string)
	dishData.RestaurantId = RestaurantIdString

	if err := c.BindJSON(&dishData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resDishData, err := u.DishUseCase.UpdateDishDetails(&dishData, &num)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't update Dish", resDishData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "dish updated succesfully", resDishData, nil)
	c.JSON(http.StatusOK, response)
}
func (u *DishHandler) DeleteDish(c *gin.Context) {
	dishid := c.Param("dishid")
	err := u.DishUseCase.DeleteDish(&dishid)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete Dish", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "dish deleted succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *DishHandler) GetAllDishesForUser(c *gin.Context) {
	dishMap, err := u.DishUseCase.GetAllDishesForUser()
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch Dishes", dishMap, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "dishes fetched succesfully", dishMap, nil)
	c.JSON(http.StatusOK, response)
}

func (u *DishHandler) FetchDishesByCategoryId(c *gin.Context) {
	categoryid := c.Param("categoryid")
	dishSlice, err := u.DishUseCase.FetchDishesByCategoryId(&categoryid)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch Dish", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "dishs fetched", dishSlice, nil)
	c.JSON(http.StatusOK, response)
}
