package usecase

import (
	"errors"
	"fmt"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	hashpassword "github.com/ashkarax/vegn-eCommerce/pkg/hash_password"
	jwttoken "github.com/ashkarax/vegn-eCommerce/pkg/jwt_token"
	"github.com/go-playground/validator/v10"
)

type restaurantUseCase struct {
	repo  interfaceRepository.IrestaurantRepo
	token config.Token
}

func NewRestaurantUseCase(repo interfaceRepository.IrestaurantRepo, token *config.Token) interfaceUseCase.IrestaurantUseCase {
	return &restaurantUseCase{repo: repo, token: *token}
}

func (r *restaurantUseCase) RestaurantSignUp(signUpData *requestmodels.RestaurantSignUpReq) (responsemodels.RestaurantSignUpRes, error) {
	var resSignUp responsemodels.RestaurantSignUpRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(signUpData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "RestaurantName":
					resSignUp.RestaurantName = "should be a valid Name. "
				case "OwnerName":
					resSignUp.OwnerName = "should be a valid Name. "
				case "Email":
					resSignUp.Email = "should be a valid email address. "
				case "Password":
					resSignUp.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					resSignUp.ConfirmPassword = "should match the first password"

				case "Description":
					resSignUp.Description = "Description should have a minimlum of 10 words "

				case "Phone":
					resSignUp.Phone = "should be a valid number include the country code also. "

				case "District":
					resSignUp.District = "Should be a valid one"
				case "Locality":
					resSignUp.Locality = "Should be a valid one"
				case "PinCode":
					resSignUp.PinCode = "Should be a number having 6 digits"
				case "GST_NO":
					resSignUp.GST_NO = "Should be a valid one"

				}
			}
			return resSignUp, err
		}
	}

	if isRestaurantExist := r.repo.IsRestaurantExist(signUpData.Phone); isRestaurantExist {
		resSignUp.RestaurantExist = "restuarant exist with this number,change phone number"
		return resSignUp, errors.New("restuarant exists, try again with another phone number")
	}
	hashedPassword := hashpassword.HashPassword(signUpData.ConfirmPassword)
	signUpData.Password = hashedPassword

	r.repo.CreateRestaurant(signUpData)
	if err != nil {
		resSignUp.Result = err.Error()
		return resSignUp, err
	}

	resSignUp.Result = "Registeration successful!!Your request is under verification by admin."
	return resSignUp, nil

}

func (r *restaurantUseCase) RestaurantLogin(loginData *requestmodels.RestaurantLoginReq) (responsemodels.RestaurantLoginRes, error) {
	var resLogin responsemodels.RestaurantLoginRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(loginData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Password":
					resLogin.Password = "Password should have four or more digit"

				case "Phone":
					resLogin.Phone = "should be a valid number include the country code also. "

					return resLogin, err
				}
			}
		}
	}

	hashedPassword, RestaurantId, status, errr := r.repo.GetHashPassAndStatus(loginData.Phone)
	if errr != nil {
		fmt.Println(errr)
		resLogin.Result = errr.Error()
		return resLogin, errr
	}
	
	passwordErr := hashpassword.CompairPassword(hashedPassword, loginData.Password)
	if passwordErr != nil {
		resLogin.Password = passwordErr.Error()
		return resLogin, passwordErr
	}

	if status == "blocked" {
		return resLogin, errors.New("restaurant is blocked by the admin")
	}

	if status == "pending" {
		return resLogin, errors.New("restaurant is on status pending,not verified by admin yet")
	}

	accessToken, accessTokenerr := jwttoken.GenerateAcessToken(r.token.RestaurantSecurityKey, RestaurantId)
	if err != nil {
		return resLogin, accessTokenerr
	}

	refreshToken, refreshTokenerr := jwttoken.GenerateRefreshToken(r.token.RestaurantSecurityKey)
	if err != nil {
		return resLogin, refreshTokenerr
	}

	resLogin.AccessToken = accessToken
	resLogin.RefreshToken = refreshToken

	return resLogin, nil
}



func (r *restaurantUseCase) RestaurantsByStatus(status string)(*[]responsemodels.RestuarntDetails,error){
	
	verResMap,err :=r.repo.GetRestaurantsByStatus(status)
	if err != nil{
		return verResMap,err
	}
return verResMap,nil

}

func (r *restaurantUseCase) ChangeRestaurantStatusById(id int,status string) error{

	err:=r.repo.ChangeRestaurantStatusById(id ,status)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}




