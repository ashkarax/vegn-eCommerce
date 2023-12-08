package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IuserRepo interface {
	IsUserExist(string) bool
	CreateUser(userData *requestmodels.UserSignUpReq)
	ChangeUserStatusActive(string) error
	GetUserId(string) (string, error)
	GetHashPassAndStatus(string) (string, string, string, error)

	GetLatestUsers() (*[]responsemodels.UserDetails, error)
	SearchUserByIdOrName(int, string) (*[]responsemodels.UserDetails, error)
	ChangeUserStatusById(int, string) error
	GetUserByStatus(string) (*[]responsemodels.UserDetails, error)

	GetUserProfile(*string) (*responsemodels.UserDetails, error) 
	EditUserDetails(*requestmodels.UserEditProf) error

	AddMoneyToWallet(*string,*float64) error

}
