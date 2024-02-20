package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
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
		fmt.Println("Error in deleting already existing user with this phonenumber with status pending")
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
	query := "INSERT INTO users (f_name,l_name, email, phone, password) VALUES($1, $2, $3, $4,$5)"
	d.DB.Exec(query, userData.FirstName, userData.LastName, userData.Email, userData.Phone, userData.Password)
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

func (d *userRepository) GetLatestUsers() (*[]responsemodels.UserDetails, error) {
	var userDataMap []responsemodels.UserDetails

	query := "SELECT * FROM users WHERE status = 'active' ORDER BY created_at DESC LIMIT 30;"
	r := d.DB.Raw(query).Scan(&userDataMap)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("Rows affected:%d", r.RowsAffected)
		return &userDataMap, errors.New(errMessage)
	}
	if r.Error != nil {
		return &userDataMap, r.Error
	}

	return &userDataMap, nil
}

func (d *userRepository) SearchUserByIdOrName(id int, name string) (*[]responsemodels.UserDetails, error) {
	var userDataMap []responsemodels.UserDetails

	r := d.DB.Raw("SELECT * FROM users WHERE f_name LIKE ? OR id = ?", "%"+name+"%", id).Scan(&userDataMap)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,Rows affected:%d", r.RowsAffected)
		return &userDataMap, errors.New(errMessage)
	}
	if r.Error != nil {
		return &userDataMap, r.Error
	}

	return &userDataMap, nil
}

func (d *userRepository) ChangeUserStatusById(id int, status string) error {
	query := "UPDATE users SET status = $1 WHERE id = $2"
	err := d.DB.Exec(query, status, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *userRepository) GetUserByStatus(status string) (*[]responsemodels.UserDetails, error) {
	var userDataMap []responsemodels.UserDetails

	query := "SELECT * FROM users WHERE status = $1"
	r := d.DB.Raw(query, status).Scan(&userDataMap)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No %s users available,Rows affected:%d", status, r.RowsAffected)
		return &userDataMap, errors.New(errMessage)
	}
	if r.Error != nil {
		return &userDataMap, r.Error
	}

	return &userDataMap, nil
}

func (d *userRepository) GetUserProfile(id *string) (*responsemodels.UserDetails, error) {
	var userData responsemodels.UserDetails

	r := d.DB.Raw("SELECT id,f_name,l_name,email,phone,status FROM users WHERE id = ?", id).Scan(&userData)

	if r.RowsAffected == 0 {
		return &userData, errors.New("no results found,Rows affected 0")
	}
	if r.Error != nil {
		return &userData, r.Error
	}

	return &userData, nil
}

func (d *userRepository) EditUserDetails(editData *requestmodels.UserEditProf) error {

	query := "UPDATE users SET f_name = ?, l_name = ?, email = ? WHERE id = ?"

	result := d.DB.Exec(query,
		editData.FName,
		editData.LName,
		editData.Email,
		editData.UserId,
	)
	if result.Error != nil {
		fmt.Println(result)
		return result.Error
	}

	return nil
}

func (d *userRepository) AddMoneyToWallet(userId *string, refundAmount *float64) error {
	query := "UPDATE users SET wallet = wallet + ? WHERE id = ?"
	err := d.DB.Exec(query, refundAmount, userId).Error
	if err != nil {
		return err
	}
	return nil

}
