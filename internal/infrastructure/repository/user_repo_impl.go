package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaceRepository.IuserRepo {
	return &userRepository{DB: DB}
}
func (d *userRepository) IsUserExist(phone string) bool {
	var userCount int

	delUncompletedUser := "DELETE FROM users WHERE phone =$1 AND status =$2"
	result := d.DB.Exec(delUncompletedUser, phone, "pending")
	if result.Error != nil {
		fmt.Println("No users wit this phno as pending")
	}

	query := "SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "deleted").Row().Scan(&userCount)
	if err != nil {
		fmt.Println("error in usercount query")
	}
	if userCount >= 1 {
		return true
	}

	return false
}

func (d *userRepository) CreateUser(userData *requestmodels.UserSignUpReq) {
	query := "INSERT INTO users (name, email, phone, password) VALUES($1, $2, $3, $4)"
	d.DB.Exec(query, userData.FirstName, userData.Email, userData.Phone, userData.Password)
}

func (d *userRepository) ChangeUserStatusActive(phone string) error {
	query := "UPDATE users SET status = 'active' WHERE phone = $1"
	result := d.DB.Exec(query, phone)
	if result.Error != nil {
		fmt.Println("", result.Error)

		return result.Error
	}

	return nil

}

func (d *userRepository) GetUserId(phone string) (string, error) {
	var userId string
	query := "SELECT id FROM users WHERE phone=$1"
	err := d.DB.Raw(query, phone).Row().Scan(&userId)
	if err != nil {
		fmt.Println("", err)
		return "", err
	}
	return userId, nil

}

func (d *userRepository) GetHashPassAndStatus(phone string) (string, string, string, error) {
	var hashedPassword, status, sellerid string

	query := "SELECT password, id, status FROM users WHERE phone=? AND status!='delete'"
	err := d.DB.Raw(query, phone).Row().Scan(&hashedPassword, &sellerid, &status)
	if err != nil {
		return "", "", "", errors.New("no user exist with the specified ph no,signup first")
	}

	return hashedPassword, sellerid, status, nil
}
