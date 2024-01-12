package main

import (
	"log"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	"github.com/ashkarax/vegn-eCommerce/internal/di"
)

// @title Go + Gin Veg*n.
// @version 1.0
// @description Online Veg Food Delivery.
// @securityDefinitions.apikey	OtpTempTokenAuth
// @in header
// @name Authorizations
// @securityDefinitions.apikey	AdminRefTokenAuth
// @in header
// @name refreshtoken
// @securityDefinitions.apikey	RestaurantAuthTokenAuth
// @in header
// @name accesstoken
// @securityDefinitions.apikey	RestaurantRefTokenAuth
// @in header
// @name refreshtoken
// @securityDefinitions.apikey	UserAuthTokenAuth
// @in header
// @name accesstoken
// @securityDefinitions.apikey	UserRefTokenAuth
// @in header
// @name refreshtoken
// @contact.name API Support
// @host localhost:8080
// @BasePath /

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("error at loading the env file using viper")
	}
	server, diErr := di.InitializeAPI(*config)
	if diErr != nil {
		log.Fatal("error for server creation")
	}

	server.Start()
}
