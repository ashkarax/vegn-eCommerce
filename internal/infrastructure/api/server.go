package server

import (
	"fmt"
	"log"

	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ashkarax/vegn-eCommerce/docs"
)

type ServerHttp struct {
	engin *gin.Engine
}

func NewServerHttp(adminHandler *handlers.AdminHandler,
	userHandler *handlers.UserHandler,
	restaurantHandler *handlers.RestaurantHandler,
	dishHandler *handlers.DishHandler,
	addressHandler *handlers.AddressHandler,
	cartHandler *handlers.CartHandler,
	orderHandler *handlers.OrderHandler,
	paymentHandler *handlers.PaymentHandler,
	couponHandler *handlers.CouponHandler,
	categoryHandler *handlers.CategoryHandler,
	JWTmiddleware *middlewares.TokenRequirements) *ServerHttp {

	engin := gin.Default()

	// load htmlpages
	engin.LoadHTMLGlob("./templates/*.html")

	
	// use ginSwagger middleware to serve the API docs
	engin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.AdminRoutes(engin.Group("/admin"), adminHandler, couponHandler, categoryHandler, JWTmiddleware)
	routes.UserRoutes(engin.Group(""), userHandler, dishHandler, addressHandler, cartHandler, orderHandler, paymentHandler, categoryHandler, JWTmiddleware)
	routes.RestaurantRoutes(engin.Group("/restaurant"), restaurantHandler, dishHandler, orderHandler, categoryHandler, JWTmiddleware)

	return &ServerHttp{engin: engin}

}

func (server *ServerHttp) Start() {
	err := server.engin.Run(":8080")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
