package handlers

import (
	"net/http"

	_ "github.com/ashkarax/vegn-eCommerce/docs"
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

//	@Summary		UserSignUp
//	@Description	User can signup using this handler
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		requestmodels.UserSignUpReq	true	"User Sign-Up Details"
//	@Success		200		{object}	responsemodels.Response
//	@Failure		400		{object}	responsemodels.Response
//	@Router			/signup/ [post]
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

//	@Summary		UserOTPVerication
//	@Description	User can verify the OTP which is generated after successful signup request.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		OtpTempTokenAuth
//	@Param			user	body		requestmodels.OtpVerification	true	"OTP generated"
//	@Success		200		{object}	responsemodels.Response
//	@Failure		400		{object}	responsemodels.Response
//	@Router			/verify/ [post]
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
	c.JSON(http.StatusOK, response)
}

//	@Summary		UserLogin
//	@Description	User can login using this handler.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		requestmodels.UserLoginReq	true	"Login Credentials"
//	@Success		200		{object}	responsemodels.Response
//	@Failure		400		{object}	responsemodels.Response
//	@Router			/login/ [post]
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

//	@Summary		GetUserProfile
//	@Description	Retrieves the profile information for a specific user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		UserAuthTokenAuth
//	@Security		UserRefTokenAuth
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Failure		500	{object}	responsemodels.Response
//	@Router			/profile [get]
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

//	@Summary		EditUserProfile
//	@Description	Edits a user's profile.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		UserAuthTokenAuth
//	@Security		UserRefTokenAuth
//	@Success		200	{object}	responsemodels.Response
//	@Failure		400	{object}	responsemodels.Response
//	@Router			/profile [patch]
func (u *UserHandler) EditUserProfile(c *gin.Context) {
	var editData requestmodels.UserEditProf

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)
	editData.UserId = userIdString

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
