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

//	@Summary		Add a new dish
//	@Description	Add a new dish for the restaurant specified in the request context
//	@Tags			RestauratDishManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Param			name					formData	string	true	"Name of the dish"
//	@Param			category_id				formData	integer	true	"Category ID of the dish"
//	@Param			description				formData	string	true	"Description of the dish"
//	@Param			cuisine_type			formData	string	true	"Cuisine type of the dish"
//	@Param			mrp						formData	number	true	"MRP (Maximum Retail Price) of the dish"
//	@Param			portion_size			formData	string	false	"Portion size of the dish"
//	@Param			dietary_information		formData	string	false	"Dietary information of the dish"
//	@Param			calories				formData	integer	false	"Calories in the dish"
//	@Param			protein					formData	integer	false	"Protein content in the dish"
//	@Param			carbohydrates			formData	integer	false	"Carbohydrates content in the dish"
//	@Param			fat						formData	integer	false	"Fat content in the dish"
//	@Param			spice_level				formData	string	false	"Spice level of the dish"	
//	@Param			allergen_information	formData	string	false	"Allergen information of the dish"
//	@Param			recommended_pairings	formData	string	false	"Recommended pairings for the dish"
//	@Param			special_features		formData	string	false	"Special features of the dish"
//	@Param			image					formData	file	true	"Image file of the dish"
//	@Param			preparation_time		formData	string	false	"Preparation time for the dish"
//	@Param			promotion_discount		formData	integer	false	"Promotion discount for the dish"
//	@Param			story_origin			formData	string	false	"Story origin of the dish"
//	@Success		200						{object}	responsemodels.Response
//	@Failure		400						{object}	responsemodels.Response
//	@Router			/restaurant/dish [post]
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

// FetchAllDishesForRestaurant is a Swaggo-documented function to fetch all dishes for a restaurant.
//	@Summary		Fetch all dishes for a restaurant
//	@Description	Fetch all dishes for the restaurant specified in the request context
//	@Tags			RestauratDishManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/restaurant/dish [get]
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

//	@Summary		FetchDishWithId
//	@Description	Retrieves details for a specific dish.
//	@Tags			Dish
//	@Accept			json
//	@Produce		json
//	@Param			dishid	path		int	true	"The ID of the dish to fetch."
//	@Success		200		{object}	responsemodels.Response
//	@Failure		400		{object}	responsemodels.Response
//	@Router			/dishes/{dishid} [get]
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

//	@Summary		Update details for a dish
//	@Description	Update details for the dish specified by dishid
//	@Tags			RestauratDishManagement
//	@Accept			json
//	@Produce		json
//	@Security		RestaurantAuthTokenAuth
//	@Security		RestaurantRefTokenAuth
//	@Param			dishid		path		int							true	"ID of the dish to update."
//	@Param			dishData	body		requestmodels.DishUpdateReq	true	"Updated dish information"
//	@Success		200			{object}	responsemodels.Response
//	@Failure		400			{object}	responsemodels.Response
//	@Router			/restaurant/dish/{dishid} [patch]
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

//	@Summary		GetAllDishesForUser
//	@Description	Retrieves all dishes.
//	@Tags			Dish
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/dishes [get]
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

//	@Summary		FetchDishesByCategoryId
//	@Description	Retrieves dishes for a specific category.
//	@Tags			Dish
//	@Accept			json
//	@Produce		json
//	@Param			categoryid	path		string	true	"The ID of the category to fetch dishes for."
//	@Success		200			{object}	responsemodels.Response
//	@Failure		400			{object}	responsemodels.Response
//	@Router			/dishes/{categoryid} [get]
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
