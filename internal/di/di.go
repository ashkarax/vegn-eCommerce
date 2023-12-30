package di

import (
	"fmt"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	server "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/api"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/db"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase"
	"github.com/ashkarax/vegn-eCommerce/pkg/aws"
	"github.com/ashkarax/vegn-eCommerce/pkg/twilio"
)

func InitializeAPI(config config.Config) (*server.ServerHttp, error) {

	DB, err := db.ConnectDatabase(config.DB)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM DI.GO")
		return nil, err
	}

	twilio.OTPServiceSetup(config.Otp)
	aws.AWSS3FileUploaderSetup(config.AwsS3)

	JWTRepository := repository.NewJWTRepo(DB)
	JWTUseCase := usecase.NewJWTUseCase(JWTRepository)
	JWTmiddleware := middlewares.NewJWTTokenMiddleware(JWTUseCase, &config.Token)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, &config.Token)
	userHandler := handlers.NewUserhandler(userUseCase)

	restaurantRepository := repository.NewRestaurantRepository(DB)
	restaurantUsecase := usecase.NewRestaurantUseCase(restaurantRepository, &config.Token)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantUsecase)

	adminRepository := repository.NewAdminRepository(DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, &config.Token)
	adminHandler := handlers.NewAdminHandler(adminUseCase, restaurantUsecase, userUseCase)

	categoryRepository := repository.NewCategoryRepo(DB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryUseCase)

	dishRepository := repository.NewDishRepo(DB)
	dishUseCase := usecase.NewDishUseCase(dishRepository, categoryRepository)
	dishHandler := handlers.NewDishHandler(dishUseCase)

	addressRepository := repository.NewAddressRepo(DB)
	addressUseCase := usecase.NewAddressUseCase(addressRepository)
	addressHandler := handlers.NewAddressHandler(addressUseCase)

	cartRepository := repository.NewCartRepo(DB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, dishRepository)
	cartHandler := handlers.NewCartHandler(cartUseCase)

	couponRepository := repository.NewCouponRepo(DB)
	couponUseCase := usecase.NewCouponUseCase(couponRepository)
	couponHandler := handlers.NewCouponHandler(couponUseCase)

	orderRepository := repository.NewOrderRepo(DB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository, dishRepository, addressRepository, userRepository, couponRepository, restaurantRepository, &config.RazorP)
	orderHandler := handlers.NewOrderHandler(orderUseCase)

	paymentUseCase := usecase.NewPaymentUsecase(orderRepository, cartRepository, &config.RazorP)
	paymentHandler := handlers.NewPaymentHandler(paymentUseCase)

	serverHttp := server.NewServerHttp(adminHandler, userHandler, restaurantHandler, dishHandler, addressHandler, cartHandler, orderHandler, paymentHandler, couponHandler, categoryHandler, JWTmiddleware)

	return serverHttp, nil

}
