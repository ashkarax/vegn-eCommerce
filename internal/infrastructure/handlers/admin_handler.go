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

type AdminHandler struct {
	AdminUseCase      interfaceUseCase.IAdminUseCase
	RestaurantUseCase interfaceUseCase.IrestaurantUseCase
	UserUsecase       interfaceUseCase.IuserUseCase
}

func NewAdminHandler(useCase interfaceUseCase.IAdminUseCase,
	restaurant interfaceUseCase.IrestaurantUseCase,
	user interfaceUseCase.IuserUseCase) *AdminHandler {
	return &AdminHandler{AdminUseCase: useCase,
		RestaurantUseCase: restaurant,
		UserUsecase:       user}
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

// From restaurant usecase
func (u *AdminHandler) VerifiedRestuarants(c *gin.Context) {

	result, err := u.RestaurantUseCase.RestaurantsByStatus("verified")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "No restaurants with status verified", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", result, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *AdminHandler) PendingRestuarants(c *gin.Context) {

	result, err := u.RestaurantUseCase.RestaurantsByStatus("pending")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "No restaurants with status verified", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", result, nil)
	c.JSON(http.StatusOK, finalReslt)
}
func (u *AdminHandler) RejectedRestuarants(c *gin.Context) {

	result, err := u.RestaurantUseCase.RestaurantsByStatus("rejected")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "No restaurants with status rejected", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", result, nil)
	c.JSON(http.StatusOK, finalReslt)
}
func (u *AdminHandler) BlockedRestuarants(c *gin.Context) {

	result, err := u.RestaurantUseCase.RestaurantsByStatus("rejected")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "No restaurants with status blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", result, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *AdminHandler) BlockRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	num, _ := strconv.Atoi(restaurantID)

	err := u.RestaurantUseCase.ChangeRestaurantStatusById(num, "blocked")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)

}

func (u *AdminHandler) VerifyRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	num, _ := strconv.Atoi(restaurantID)

	err := u.RestaurantUseCase.ChangeRestaurantStatusById(num, "verified")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)

}

func (u *AdminHandler) RejectRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	num, _ := strconv.Atoi(restaurantID)

	err := u.RestaurantUseCase.ChangeRestaurantStatusById(num, "rejected")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)

}
func (u *AdminHandler) UnBlockRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	num, _ := strconv.Atoi(restaurantID)

	err := u.RestaurantUseCase.ChangeRestaurantStatusById(num, "verified")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)

}
func (u *AdminHandler) DeleteRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	num, _ := strconv.Atoi(restaurantID)

	err := u.RestaurantUseCase.ChangeRestaurantStatusById(num, "deleted")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)

}

func (u *AdminHandler) LatestUsers(c *gin.Context) {

	usersMap, err := u.UserUsecase.GetLatestUsers()
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "something went wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", usersMap, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *AdminHandler) SearchUser(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	if id == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide either ID or name"})
		return
	}
	num, _ := strconv.Atoi(id)
	usersMap, err := u.UserUsecase.SearchUserByIdOrName(num, name)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "something went wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", usersMap, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *AdminHandler) BlockUser(c *gin.Context) {
	userId := c.Param("id")
	num, _ := strconv.Atoi(userId)

	err := u.UserUsecase.ChangeUserStatusById(num, "blocked")
	if err != nil {
		result := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, result)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *AdminHandler) UnBlockUser(c *gin.Context) {
	userId := c.Param("id")
	num, _ := strconv.Atoi(userId)

	err := u.UserUsecase.ChangeUserStatusById(num, "active")
	if err != nil {
		result := responsemodels.Responses(http.StatusBadRequest, "Error in changing status", nil, err.Error())
		c.JSON(http.StatusBadRequest, result)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully updated status", " ", nil)
	c.JSON(http.StatusOK, finalReslt)

}

func (u *AdminHandler) BlockedUsers(c *gin.Context) {
	blockedUsersmap, err := u.UserUsecase.UserByStatus("blocked")
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "No users with status blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", blockedUsersmap, nil)
	c.JSON(http.StatusOK, finalReslt)

}
