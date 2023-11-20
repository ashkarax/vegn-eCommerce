package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup,user *handlers.UserHandler){
	engin.POST("/signup",user.UserSignUp)
	engin.POST("/verify",user.UserOTPVerication)
	engin.POST("/login",user.UserLogin)
}