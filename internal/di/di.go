package di

import (
	"fmt"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	server "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/api"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/db"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase"
	"github.com/ashkarax/vegn-eCommerce/pkg/twilio"
)

func InitializeAPI(config config.Config) (*server.ServerHttp, error) {

	DB, err := db.ConnectDatabase(config.DB)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM DI.GO")
		return nil, err
	}
	twilio.OTPServiceSetup(config.Otp)

	adminRepository := repository.NewAdminRepository(DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, &config.Token)
	adminHandler := handlers.NewAdminHandler(adminUseCase)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, &config.Token)
	userHandler := handlers.NewUserhandler(userUseCase)

	sellerRepository:=repository.NewSellerRepository(DB)
	sellerUsecase := usecase.NewSellerUseCase(sellerRepository,&config.Token)
	sellerHandler := handlers.NewSellerHandler(sellerUsecase)

	serverHttp := server.NewServerHttp(adminHandler, userHandler, sellerHandler)

	return serverHttp, nil

}
