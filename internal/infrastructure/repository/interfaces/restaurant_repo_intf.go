package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IrestaurantRepo interface {
	IsRestaurantExist(string) bool
	CreateRestaurant(*requestmodels.RestaurantSignUpReq) error
	GetHashPassAndStatus(string) (string, string, string, error)
	GetRestaurantsByStatus(string) (*[]responsemodels.RestuarntDetails, error)
	ChangeRestaurantStatusById(id int, status string) error

	AddMoneyToCodWallet(*float64, *string) error
	AddMoneyToAdminCredit(*float64, *string) error
}
