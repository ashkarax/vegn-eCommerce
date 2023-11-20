package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IuserUseCase interface {
	UserSignUp(*requestmodels.UserSignUpReq) (responsemodels.SignupData, error)
	VerifyOtp(*requestmodels.OtpVerification, string) (responsemodels.OtpVerifResult,error)
	UserLogin(*requestmodels.UserLoginReq) (responsemodels.UserLoginRes,error)
}
