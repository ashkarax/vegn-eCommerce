package server

import (
	"fmt"
	"log"

	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engin *gin.Engine
}

func NewServerHttp(admin *handlers.AdminHandler,user *handlers.UserHandler ,seller *handlers.SellerHandler) *ServerHttp {

	engin := gin.Default()

	// engin := gin.New()
	// engin.Use(gin.Logger())

	routes.AdminRoutes(engin.Group("/admin"), admin)
	routes.UserRoutes(engin.Group(""),user)
	routes.SellerRoutes(engin.Group("/restaurant"),seller)

	return &ServerHttp{engin: engin}

}

func (server *ServerHttp) Start() {
	err := server.engin.Run(":8080")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
