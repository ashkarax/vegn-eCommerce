package usecase

import (
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/go-playground/validator/v10"
)

type AddressUseCase struct {
	Repo interfaceRepository.IAddressRepo
}

func NewAddressUseCase(addressRepo interfaceRepository.IAddressRepo) interfaceUseCase.IAddressUseCase {
	return &AddressUseCase{Repo: addressRepo}
}
func (r *AddressUseCase) AddNewAddress(addressData *requestmodels.AddressReq) (*responsemodels.AddressRes, error) {
	var resAddressData responsemodels.AddressRes
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(addressData)
	if err != nil {


//need to do proper validation

		fmt.Println(err)
		return &resAddressData, err
	}
	insertErr := r.Repo.AddNewAddress(addressData)
	if insertErr != nil {
		fmt.Println(insertErr)
		return &resAddressData, insertErr
	}
	return &resAddressData, nil

	
}
func (r *AddressUseCase)  EditAddress(addressData *requestmodels.AddressReq) (*responsemodels.AddressRes,error){
	var resAddressData responsemodels.AddressRes
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(addressData)
	if err != nil {


//need to do proper validation

		fmt.Println(err)
		return &resAddressData, err
	}
	insertErr := r.Repo.EditAddress(addressData)
	if insertErr != nil {
		fmt.Println(insertErr)
		return &resAddressData, insertErr
	}
	return &resAddressData, nil
}
func (r *AddressUseCase) GetAllAddresses(userId *string)(*[]responsemodels.AddressRes,error){
	addressMap,err := r.Repo.GetUserAddresses(userId)
	if err != nil {
		fmt.Println(err)
		return addressMap, err
	}
	return addressMap, nil
}


