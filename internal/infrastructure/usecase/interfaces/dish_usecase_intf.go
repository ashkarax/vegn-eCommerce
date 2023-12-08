package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IDishUseCase interface {
	NewDish(*requestmodels.DishReq) (*responsemodels.DishRes, error)
	FetchAllDishesForRestaurant(*int) (*[]responsemodels.DishRes, error)
	DishById(*int) (*responsemodels.DishRes, error)
	UpdateDishDetails(*requestmodels.DishUpdateReq, *int) (*responsemodels.DishRes, error)
	DeleteDish(*string) error
	GetAllDishesForUser() (*[]responsemodels.DishRes, error)

	FetchDishesByCategoryId(*string) (*[]responsemodels.DishRes, error)
}
