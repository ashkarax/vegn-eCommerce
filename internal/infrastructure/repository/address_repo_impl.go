package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type AddressRepo struct {
	DB *gorm.DB
}

func NewAddressRepo(DB *gorm.DB) interfaceRepository.IAddressRepo {
	return &AddressRepo{DB: DB}
}

func (d *AddressRepo) AddNewAddress(addData *requestmodels.AddressReq) error {
	query := "INSERT INTO addresses (user_id, line1, street, city, state, postal_code, country, phone, alternate_phone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result := d.DB.Exec(query,
		addData.UserID,
		addData.Line1,
		addData.Street,
		addData.City,
		addData.State,
		addData.PostalCode,
		addData.Country,
		addData.Phone,
		addData.AlternatePhone,
	).Error
	if result != nil {
		fmt.Println(result)
		return result
	}
	return nil
}

func (d *AddressRepo) EditAddress(addData *requestmodels.AddressReq) error {
	query := "UPDATE addresses SET line1 = ?, street = ?, city = ?, state = ?, postal_code = ?, country = ?, phone = ?, alternate_phone = ? WHERE user_id = ? AND id = ?"
	result := d.DB.Exec(query,
		addData.Line1,
		addData.Street,
		addData.City,
		addData.State,
		addData.PostalCode,
		addData.Country,
		addData.Phone,
		addData.AlternatePhone,
		addData.UserID,
		addData.AddressId,
	)
	if result.Error != nil {
		fmt.Println(result)
		return result.Error
	}
	if result.RowsAffected == 0 {
		errMessage := fmt.Sprintf("Address does not match with User,Rows affected:%d", result.RowsAffected)
		return errors.New(errMessage)

	}
	return nil
}
func (d *AddressRepo) GetUserAddresses(userId *string) (*[]responsemodels.AddressRes, error) {
	var resAddrMap []responsemodels.AddressRes
	result := d.DB.Raw("SELECT * FROM addresses WHERE user_id=?", userId).Scan(&resAddrMap)
	if result.Error != nil {
		fmt.Println(result)
		return &resAddrMap, result.Error
	}
	if result.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No Address on this user-Id,Rows affected:%d", result.RowsAffected)
		return &resAddrMap, errors.New(errMessage)

	}
	return &resAddrMap, nil

}

func (d *AddressRepo)  VerifyAddress(orderDetails *requestmodels.OrderDetails) (int,error) {
var count int 
query:="SELECT COUNT(*) FROM addresses WHERE user_id=? AND id=?"
err:=d.DB.Raw(query,orderDetails.UserID,orderDetails.AddressID).Scan(&count).Error
if err !=nil{
	return 0,err
}
if count == 0{
	return 0,nil
}

return count,nil

}

