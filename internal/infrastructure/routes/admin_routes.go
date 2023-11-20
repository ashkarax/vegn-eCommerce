package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handlers.AdminHandler) {
	engin.POST("/login", admin.AdminLogin)

}
