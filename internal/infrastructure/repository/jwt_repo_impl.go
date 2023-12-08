package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	"gorm.io/gorm"
)

type JWTRepo struct {
	DB *gorm.DB
}

func NewJWTRepo(DB *gorm.DB) interfaceRepository.IJWTRepo {
	return &JWTRepo{DB: DB}
}

func (d *JWTRepo) GetRestStatForGeneratingAccessToken(restaurantid *int) (*string,error) {
var restaurantCurrentStatus string
query := "SELECT status from restaurants WHERE id=?"
result := d.DB.Raw(query, restaurantid).Scan(&restaurantCurrentStatus)
	
	if result.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,No user with this id=%d found in db",*restaurantid)
		return &restaurantCurrentStatus,errors.New(errMessage)
	}
	if result.Error != nil {
		return &restaurantCurrentStatus,result.Error
	}

	return &restaurantCurrentStatus,nil
}

func (d *JWTRepo) GetUserStatForGeneratingAccessToken(userId *string) (*string,error){
	var userCurrentStatus string
query := "SELECT status from users WHERE id=?"
result := d.DB.Raw(query, userId).Scan(&userCurrentStatus)
	
	if result.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,No user with this id=%s found in db",*userId)
		return &userCurrentStatus,errors.New(errMessage)
	}
	if result.Error != nil {
		return &userCurrentStatus,result.Error
	}

	return &userCurrentStatus,nil
}
