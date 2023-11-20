package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IAdminUseCase interface {
	AdminLogin(*requestmodels.AdminLoginData) (*responsemodels.AdminLoginRes, error)
}
