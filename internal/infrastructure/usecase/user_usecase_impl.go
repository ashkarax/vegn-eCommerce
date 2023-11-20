package usecase

import (
	"errors"
	"github.com/ashkarax/vegn-eCommerce/internal/config"
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	hashpassword "github.com/ashkarax/vegn-eCommerce/pkg/hash_password"
	jwttoken "github.com/ashkarax/vegn-eCommerce/pkg/jwt_token"
	"github.com/ashkarax/vegn-eCommerce/pkg/twilio"
	"github.com/go-playground/validator/v10"
)

type UserUsecase struct {
	repo             interfaceRepository.IuserRepo
	tokenSecurityKey *config.Token
}

func NewUserUseCase(userRepository interfaceRepository.IuserRepo, tokenSecurityKey *config.Token) interfaceUseCase.IuserUseCase {
	return &UserUsecase{repo: userRepository, tokenSecurityKey: tokenSecurityKey}
}

func (r *UserUsecase) UserSignUp(userData *requestmodels.UserSignUpReq) (responsemodels.SignupData, error) {

	var resSignUp responsemodels.SignupData

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "FirstName":
					resSignUp.Name = "should be a valid Name. "
				case "Phone":
					resSignUp.Phone = "should include the country code also. "
				case "Email":
					resSignUp.Email = "should be a valid email address. "
				case "Password":
					resSignUp.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					resSignUp.ConfirmPassword = "should match the first password"
				}
			}
		}
		return resSignUp, errors.New("did't fullfill the signup requirement ")
	}

	if isUserExist := r.repo.IsUserExist(userData.Phone); isUserExist {
		resSignUp.IsUserExist = "User Exist ,change phone number"
		return resSignUp, errors.New("user exists, try again with another phone number")
	}

	client := twilio.TwilioClient()
	_, otpSndErr := twilio.SendOTP(userData.Phone, client)
	if otpSndErr != nil {
		resSignUp.OTP = otpSndErr.Error()
		return resSignUp, errors.New("error at attempt for creating OTP")
	}

	hashedPassword := hashpassword.HashPassword(userData.ConfirmPassword)
	userData.Password = hashedPassword

	r.repo.CreateUser(userData)
	tempToken, err := jwttoken.TempTokenForOtpVerification(r.tokenSecurityKey.TempVerificationKey, userData.Phone)
	if err != nil {
		resSignUp.Token = "error creating temp token for otp verification"
		return resSignUp, errors.New("error creating token")
	}

	resSignUp.Token = tempToken

	return resSignUp, nil

}

func (r *UserUsecase) VerifyOtp(otpData *requestmodels.OtpVerification, TempVerificationToken string) (responsemodels.OtpVerifResult, error) {
	var otpveriRes responsemodels.OtpVerifResult

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(otpData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					otpData.Otp = "otp should be a 4 digit number"
				}
			}
		}
		return otpveriRes, errors.New("did't fullfill the login requirement ")
	}
	phone, unbindErr := jwttoken.UnbindPhoneFromClaim(TempVerificationToken, r.tokenSecurityKey.TempVerificationKey)
	if unbindErr != nil {
		otpveriRes.Token = "invalid token"
		return otpveriRes, unbindErr
	}

	client := twilio.TwilioClient()
	otpSndErr := twilio.VerifyOtp(phone, otpData.Otp, client)
	if otpSndErr != nil {
		otpveriRes.Otp = otpSndErr.Error()
		return otpveriRes, errors.New("OTP does not match with your phone number")
	}

	changeStatErr := r.repo.ChangeUserStatusActive(phone)
	if changeStatErr != nil {
		return otpveriRes, changeStatErr
	}

	userId, fetchErr := r.repo.GetUserId(phone)
	if fetchErr != nil {
		return otpveriRes, fetchErr
	}

	accessToken, aTokenErr := jwttoken.GenerateAcessToken(r.tokenSecurityKey.UserSecurityKey, userId)
	if aTokenErr != nil {
		otpveriRes.AccessToken = aTokenErr.Error()
		return otpveriRes, aTokenErr
	}
	refreshToken, rTokenErr := jwttoken.GenerateRefreshToken(r.tokenSecurityKey.UserSecurityKey)
	if rTokenErr != nil {
		otpveriRes.RefreshToken = rTokenErr.Error()
		return otpveriRes, rTokenErr
	}
	otpveriRes.Otp = "verified"
	otpveriRes.AccessToken = accessToken
	otpveriRes.RefreshToken = refreshToken

	return otpveriRes, nil
}

func (r *UserUsecase) UserLogin(loginData *requestmodels.UserLoginReq) (responsemodels.UserLoginRes, error) {
	var resLogin responsemodels.UserLoginRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(loginData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Phone":
					resLogin.Phone = "should be 10 digits excluding the country "
				case "Password":
					resLogin.Password = "Password should have four or more digit"
				}
			}
			return resLogin, errors.New("did't fullfill the login requirement ")
		}
	}

	hashedPassword, userId, status, errr := r.repo.GetHashPassAndStatus(loginData.Phone)
	if errr != nil {
		return resLogin, errr
	}

	passwordErr := hashpassword.CompairPassword(hashedPassword, loginData.Password)
	if passwordErr != nil {
		return resLogin, passwordErr
	}

	if status == "blocked" {
		return resLogin, errors.New("user is blocked by the admin")
	}

	if status == "pending" {
		return resLogin, errors.New("user is on status pending,OTP not verified")
	}

	accessToken, accessTokenerr := jwttoken.GenerateAcessToken(r.tokenSecurityKey.UserSecurityKey, userId)
	if err != nil {
		return resLogin, accessTokenerr
	}

	refreshToken, refreshTokenerr := jwttoken.GenerateRefreshToken(r.tokenSecurityKey.UserSecurityKey)
	if err != nil {
		return resLogin, refreshTokenerr
	}

	resLogin.AccessToken = accessToken
	resLogin.RefreshToken = refreshToken
	return resLogin, nil

}
