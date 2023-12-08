package interfaceUseCase

import responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"


type ICartUseCase interface{
	AddToCart(*string,*string) error
	GetCartDetails(*string) (*responsemodels.CartDetailsResp,error)
	DeleteFromCart(*string,*string) error
}