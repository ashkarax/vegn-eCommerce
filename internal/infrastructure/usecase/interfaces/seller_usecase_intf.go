package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IsellerUseCase interface{
	SellerSignUp (*requestmodels.SellerSignUpReq) (responsemodels.SellerSignUpRes,error)
	SellerLogin (*requestmodels.SellerLoginReq) (responsemodels.SellerLoginRes,error)
}