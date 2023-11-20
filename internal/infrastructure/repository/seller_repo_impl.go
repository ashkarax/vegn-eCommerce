package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	"gorm.io/gorm"
)

type SellerUseCase struct {
	DB *gorm.DB
}

func NewSellerRepository(DB *gorm.DB) interfaceRepository.IsellerRepo {
	return &SellerUseCase{DB: DB}
}

func (d *SellerUseCase) IsSellerExist(phone string) bool {
	var sellerCount int

	delUncompletedUser := "DELETE FROM sellers WHERE contact_no =$1 AND status =$2"
	result := d.DB.Exec(delUncompletedUser, phone, "pending")
	if result.Error != nil {
		fmt.Println("No restaurants with this phno having status as pending")
	}

	query := "SELECT COUNT(*) FROM sellers WHERE contact_no=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "deleted").Row().Scan(&sellerCount)
	if err != nil {
		fmt.Println("error in restaurantCount query")
		return false
	}
	if sellerCount >= 1 {
		return true
	}

	return false
}
func (d *SellerUseCase) CreateSeller(signUpData *requestmodels.SellerSignUpReq) error {

	query := "INSERT INTO sellers (restaurant_name, owner_name, email, password, description, contact_no, district, locality, gst_no, pin_code) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	result:=d.DB.Exec(query, signUpData.RestaurantName, signUpData.OwnerName,  signUpData.Email, signUpData.Password, signUpData.Description, signUpData.ContactNo,signUpData.District, signUpData.Locality, signUpData.GST_NO, signUpData.PinCode)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *SellerUseCase) GetHashPassAndStatus(phone string) (string,string,string,error){
	var hashedPassword,status,sellerid string

	query := "SELECT password, id, status FROM sellers WHERE contact_no=? AND status!='delete'"
	err := d.DB.Raw(query, phone).Row().Scan(&hashedPassword, &sellerid, &status)
	if err != nil {
		return "", "", "", errors.New("no restuaurant exist with the specified ph no")
	}

	return  hashedPassword,sellerid,status,nil
}
