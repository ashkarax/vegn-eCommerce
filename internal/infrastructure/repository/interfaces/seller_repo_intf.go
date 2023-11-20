package interfaceRepository

import requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"

type IsellerRepo interface{
	IsSellerExist(string) bool
	CreateSeller(*requestmodels.SellerSignUpReq) error
	GetHashPassAndStatus(string) (string,string,string,error)
}