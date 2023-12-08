package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type RestaurantRepo struct {
	DB *gorm.DB
}

func NewRestaurantRepository(DB *gorm.DB) interfaceRepository.IrestaurantRepo {
	return &RestaurantRepo{DB: DB}
}

func (d *RestaurantRepo) IsRestaurantExist(phone string) bool {
	var restaurantCount int

	delUncompletedUser := "DELETE FROM restaurants WHERE phone =$1 AND status =$2"
	result := d.DB.Exec(delUncompletedUser, phone, "pending")
	if result.Error != nil {
		fmt.Println("No restaurants with this phno having status as pending")
	}

	query := "SELECT COUNT(*) FROM restaurants WHERE phone=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "deleted").Row().Scan(&restaurantCount)
	if err != nil {
		fmt.Println("error in restaurantCount query")
		return false
	}
	if restaurantCount >= 1 {
		return true
	}

	return false
}
func (d *RestaurantRepo) CreateRestaurant(signUpData *requestmodels.RestaurantSignUpReq) error {

	query := "INSERT INTO restaurants (restaurant_name, owner_name, email, password, description, phone, district, locality, gst_no, pin_code) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	result := d.DB.Exec(query, signUpData.RestaurantName, signUpData.OwnerName, signUpData.Email, signUpData.Password, signUpData.Description, signUpData.Phone, signUpData.District, signUpData.Locality, signUpData.GST_NO, signUpData.PinCode)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *RestaurantRepo) GetHashPassAndStatus(phone string) (string, string, string, error) {
	var hashedPassword, status, restaurantid string

	query := "SELECT password, id, status FROM restaurants WHERE phone=? AND status!='delete'"
	err := d.DB.Raw(query, phone).Row().Scan(&hashedPassword, &restaurantid, &status)
	if err != nil {
		return "", "", "", errors.New("no restuaurant exist with the specified ph no")
	}

	return hashedPassword, restaurantid, status, nil
}

func (d *RestaurantRepo) GetRestaurantsByStatus(status string) (*[]responsemodels.RestuarntDetails, error) {
	var restaurants []responsemodels.RestuarntDetails

	query := "SELECT * FROM restaurants WHERE status=?"
	r := d.DB.Raw(query, status).Scan(&restaurants)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("Rows affected:%d", r.RowsAffected)
		return &restaurants, errors.New(errMessage)
	}
	if r.Error != nil {
		return &restaurants, r.Error
	}

	return &restaurants, nil
}

func (d *RestaurantRepo) ChangeRestaurantStatusById(id int, status string) error {
	query := "UPDATE restaurants SET status = $1 WHERE id = $2"
	err := d.DB.Exec(query, status, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *RestaurantRepo) AddMoneyToCodWallet(codAmount *float64, restId *string) error {
	query := "UPDATE restaurants SET cod_wallet=cod_wallet + ? WHERE id = $2"
	r := d.DB.Exec(query, codAmount, restId)
	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("Rows affected:%d", r.RowsAffected)
		return errors.New(errMessage)
	}
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (d *RestaurantRepo) AddMoneyToAdminCredit(codAmount *float64, restId *string) error {
	query := "UPDATE restaurants SET admin_credit=admin_credit + ? WHERE id = $2"
	r := d.DB.Exec(query, codAmount, restId)
	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("Rows affected:%d", r.RowsAffected)
		return errors.New(errMessage)
	}
	if r.Error != nil {
		return r.Error
	}

	return nil
}


