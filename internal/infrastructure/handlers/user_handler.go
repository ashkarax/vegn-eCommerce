package handlers

import (
	"net/http"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaceUseCase.IuserUseCase
}

func NewUserhandler(userUseCase interfaceUseCase.IuserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (u *UserHandler) UserSignUp(c *gin.Context) {
	var userSignupData requestmodels.UserSignUpReq
	if err := c.BindJSON(&userSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resSignup, err := u.userUseCase.UserSignUp(&userSignupData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "signup failed", resSignup, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "signup success", resSignup, nil)
	c.JSON(http.StatusOK, response)

}

func (u *UserHandler) UserOTPVerication(c *gin.Context) {

	var otpData requestmodels.OtpVerification
	token := c.Request.Header.Get("Authorizations")

	if err := c.BindJSON(&otpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, otpVerifErr := u.userUseCase.VerifyOtp(&otpData, token)
	if otpVerifErr != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "OTP verification failed", result, otpVerifErr.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusAccepted, "OTP verification success", result, nil)
	c.JSON(http.StatusBadRequest, response)
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	var loginData requestmodels.UserLoginReq

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resLogin, err := u.userUseCase.UserLogin(&loginData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "login failed", resLogin, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "login success", resLogin, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetUserProfile(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	usersMap, err := u.userUseCase.UserProfile(&userIdString)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "something went wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", usersMap, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *UserHandler) EditUserProfile(c *gin.Context) {
	var editData requestmodels.UserEditProf

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	editData.UserId=userIdString

	if err := c.BindJSON(&editData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resEdit, err := u.userUseCase.EditUserData(&editData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to update user profile", resEdit, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "prrofile updated succesfully", resEdit, nil)
	c.JSON(http.StatusOK, response)

}


