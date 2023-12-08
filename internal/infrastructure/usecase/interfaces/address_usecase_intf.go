package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IAddressUseCase interface {
	AddNewAddress(*requestmodels.AddressReq) (*responsemodels.AddressRes,error)
	EditAddress(*requestmodels.AddressReq) (*responsemodels.AddressRes,error)
	GetAllAddresses(*string)(*[]responsemodels.AddressRes,error)
}
