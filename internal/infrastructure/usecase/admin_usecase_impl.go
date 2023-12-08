package usecase

import (
	"errors"
	"fmt"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	hashpassword "github.com/ashkarax/vegn-eCommerce/pkg/hash_password"
	jwttoken "github.com/ashkarax/vegn-eCommerce/pkg/jwt_token"
	"github.com/go-playground/validator/v10"
)

type adminUseCase struct {
	repo             interfaceRepository.IAdminRepository
	tokenSecurityKey config.Token
}

func NewAdminUseCase(adminRepository interfaceRepository.IAdminRepository, key *config.Token) interfaceUseCase.IAdminUseCase {
	return &adminUseCase{repo: adminRepository, tokenSecurityKey: *key}
}

func (r *adminUseCase) AdminLogin(adminData *requestmodels.AdminLoginData) (*responsemodels.AdminLoginRes, error) {
	var adminLoginRes responsemodels.AdminLoginRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(adminData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					adminLoginRes.Email = "should be a valid email address. "
				case "Password":
					adminLoginRes.Password = "pshould have four or more digit"
				}
			}
		}
		return &adminLoginRes, errors.New("did't fullfill the login requirement ")
	}

	HashedPassword, err := r.repo.GetPassword(adminData.Email)
	if err != nil {
		return nil, err
	}

	err = hashpassword.CompairPassword(HashedPassword, adminData.Password)
	if err != nil {
		adminLoginRes.Password = "wrong password"
		return &adminLoginRes, err
	}

	token, err := jwttoken.GenerateRefreshToken(r.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		fmt.Println("error while creating refresh token for admin")
	}

	adminLoginRes.Token = token
	return &adminLoginRes, nil

}


