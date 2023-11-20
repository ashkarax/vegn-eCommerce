package interfaceRepository

import requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"

type IuserRepo interface{
	IsUserExist(string) bool 
	CreateUser(userData *requestmodels.UserSignUpReq)
	ChangeUserStatusActive(string) error
	GetUserId(string) (string,error)
	GetHashPassAndStatus(string)(string,string,string,error)
}