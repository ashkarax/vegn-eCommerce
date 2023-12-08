package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IrestaurantUseCase interface{
	RestaurantSignUp (*requestmodels.RestaurantSignUpReq) (responsemodels.RestaurantSignUpRes,error)
	RestaurantLogin (*requestmodels.RestaurantLoginReq) (responsemodels.RestaurantLoginRes,error)

	RestaurantsByStatus(string)(*[]responsemodels.RestuarntDetails,error)
	ChangeRestaurantStatusById(id int,status string) error

}