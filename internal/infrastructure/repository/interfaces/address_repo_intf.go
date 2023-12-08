package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type IAddressRepo interface{
	AddNewAddress(*requestmodels.AddressReq) error
	EditAddress(*requestmodels.AddressReq) error
	GetUserAddresses(*string)(*[]responsemodels.AddressRes,error)
	VerifyAddress(*requestmodels.OrderDetails) (int,error)

}

