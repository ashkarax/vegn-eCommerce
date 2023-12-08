package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IuserUseCase interface {
	UserSignUp(*requestmodels.UserSignUpReq) (responsemodels.SignupData, error)
	VerifyOtp(*requestmodels.OtpVerification, string) (responsemodels.OtpVerifResult,error)
	UserLogin(*requestmodels.UserLoginReq) (responsemodels.UserLoginRes,error)

	GetLatestUsers()(* []responsemodels.UserDetails,error)
	SearchUserByIdOrName(int,string) (* []responsemodels.UserDetails,error)
	UserByStatus(string) (*[]responsemodels.UserDetails, error)
	ChangeUserStatusById(int, string) error

	UserProfile(*string) (*responsemodels.UserDetails, error) 
	EditUserData(*requestmodels.UserEditProf)(*responsemodels.UserDetails,error)


}
