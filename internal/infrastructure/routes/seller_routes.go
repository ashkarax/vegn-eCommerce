package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handlers.SellerHandler) {
	engin.POST("/signup",seller.SellerSignUp)
	engin.POST("/login",seller.SellerLogin)

}
