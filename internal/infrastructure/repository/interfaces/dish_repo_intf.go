package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IDishRepo interface {
	AddNewDish(*requestmodels.DishReq) error
	FetchAllDishesForARestaurant(*int) (*[]responsemodels.DishRes, error)
	FetchDishById(*int) (*responsemodels.DishRes, error)
	UpdateDish(*requestmodels.DishUpdateReq, *int) error
	DeleteDishById(*string) error
	GetAllDishesForUser() (*[]responsemodels.DishRes, error)

	ReturnRestaurantIdofDish(*string) (*string, error)
	DecrementDishQuantity(*string, *uint) error

	IncrementDishQuantity(*string, *uint) error

	FetchDishesByCategoryId(*string) (*[]responsemodels.DishRes, error)
}
