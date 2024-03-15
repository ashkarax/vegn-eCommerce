package repository

import (
	"errors"
	"fmt"
	"strconv"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type DishRepo struct {
	DB *gorm.DB
}

func NewDishRepo(DB *gorm.DB) interfaceRepository.IDishRepo {
	return &DishRepo{DB: DB}
}

func (d *DishRepo) AddNewDish(dishData *requestmodels.DishReq) error {

	query := "INSERT INTO dishes (restaurant_id, name, description, cuisine_type, mrp, portion_size, dietary_information, calories,protein,carbohydrates,fat, spice_level, allergen_information, recommended_pairings, special_features, image_url1,image_url2,image_url3, preparation_time, promotion_discount,price, story_origin) VALUES (?)"
	values := []interface{}{
		dishData.RestaurantId,
		dishData.Name,
		dishData.Description,
		dishData.CuisineType,
		dishData.MRP,
		dishData.PortionSize,
		dishData.DietaryInformation,
		dishData.Calories,
		dishData.Protein,
		dishData.Carbohydrates,
		dishData.Fat,
		dishData.SpiceLevel,
		dishData.AllergenInformation,
		dishData.RecommendedPairings,
		dishData.SpecialFeatures,
		dishData.ImageURL1,
		dishData.ImageURL2,
		dishData.ImageURL3,
		dishData.PreparationTime,
		dishData.PromotionDiscount,
		dishData.Price,
		dishData.StoryOrigin,
	}
	err := d.DB.Exec(query, values).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (d *DishRepo) FetchAllDishesForARestaurant(restaurant_id *int) (*[]responsemodels.DishRes, error) {
	var resMap []responsemodels.DishRes
	r := d.DB.Raw("SELECT * FROM dishes WHERE restaurant_id = ?", restaurant_id).Scan(&resMap)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,Rows affected:%d", r.RowsAffected)
		return &resMap, errors.New(errMessage)
	}
	if r.Error != nil {
		return &resMap, r.Error
	}

	return &resMap, nil

}
func (d *DishRepo) FetchDishById(dishId *int) (*responsemodels.DishRes, error) {
	var resDish responsemodels.DishRes
	r := d.DB.Raw("SELECT * FROM dishes WHERE id = ?", dishId).Scan(&resDish)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,Rows affected:%d", r.RowsAffected)
		return &resDish, errors.New(errMessage)
	}
	if r.Error != nil {
		return &resDish, r.Error
	}

	return &resDish, nil
}

func (d *DishRepo) UpdateDish(dishData *requestmodels.DishUpdateReq, dishId *int) error {
	query := "UPDATE dishes SET name = ?, category_id = ?,description = ?, cuisine_type = ?, mrp = ?, portion_size = ?, dietary_information = ?,calories = ?, protein = ?, carbohydrates = ?, fat = ?, spice_level = ?, allergen_information = ?,   recommended_pairings = ?, special_features = ?,   preparation_time = ?, promotion_discount = ?,price=?, story_origin = ?, availability = ?, remaining_quantity = ? WHERE restaurant_id = ? AND id = ?"
	result := d.DB.Exec(query,
		dishData.Name,
		dishData.CategoryId,
		dishData.Description,
		dishData.CuisineType,
		dishData.MRP,
		dishData.PortionSize,
		dishData.DietaryInformation,
		dishData.Calories,
		dishData.Protein,
		dishData.Carbohydrates,
		dishData.Fat,
		dishData.SpiceLevel,
		dishData.AllergenInformation,
		dishData.RecommendedPairings,
		dishData.SpecialFeatures,
		dishData.PreparationTime,
		dishData.PromotionDiscount,
		dishData.Price,
		dishData.StoryOrigin,
		dishData.Availability,
		dishData.RemainingQuantity,
		dishData.RestaurantId,
		dishId,
	)

	if result.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,Rows affected:%d,enter a dishid which is particularly under this restaurant", result.RowsAffected)
		return errors.New(errMessage)
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (d *DishRepo) DeleteDishById(dishId *string) error {
	restaurantId := 4
	num, _ := strconv.Atoi(*dishId)
	query := "DELETE from dishes WHERE id=? AND restaurant_id=? AND availability=?"
	result := d.DB.Exec(query, num, restaurantId, false)
	if result.RowsAffected == 0 {
		errMessage := fmt.Sprint("No results found,Dish's availability status may be true(turn it off and try again) or enter a dishid which is particularly under this restaurant")
		return errors.New(errMessage)
	}
	if result.Error != nil {
		return result.Error
	}

	return nil

}
func (d *DishRepo) GetAllDishesForUser() (*[]responsemodels.DishRes, error) {
	var resMap []responsemodels.DishRes

	r := d.DB.Raw("SELECT * FROM dishes  LEFT JOIN category_offers  ON category_offers.category_id = dishes.category_id AND category_offers.restaurant_id = dishes.restaurant_id AND category_offers.offer_status = 'active' AND category_offers.end_date >= now() WHERE dishes.availability=true AND dishes.remaining_quantity>0;").Scan(&resMap)

	if r.Error != nil {
		return &resMap, r.Error
	}

	return &resMap, nil
}

func (d *DishRepo) ReturnRestaurantIdofDish(dishid *string) (*string, error) {
	var restId string

	r := d.DB.Raw("SELECT restaurant_id FROM dishes WHERE id = ?", dishid).Scan(&restId)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,Rows affected:%d", r.RowsAffected)
		return &restId, errors.New(errMessage)
	}
	if r.Error != nil {
		return &restId, r.Error
	}

	return &restId, nil
}

func (d *DishRepo) DecrementDishQuantity(DishID *string, Quantity *uint) error {
	query := "UPDATE dishes SET remaining_quantity = remaining_quantity - ? WHERE id = ?"
	err := d.DB.Exec(query, Quantity, DishID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepo) IncrementDishQuantity(DishID *string, Quantity *uint) error {
	query := "UPDATE dishes SET remaining_quantity = remaining_quantity + ? WHERE id = ?"
	err := d.DB.Exec(query, Quantity, DishID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepo) FetchDishesByCategoryId(id *string) (*[]responsemodels.DishRes, error) {
	var resMap []responsemodels.DishRes
	r := d.DB.Raw("SELECT * FROM dishes WHERE availability = ? AND remaining_quantity >= ? AND category_id=?", true, 1, id).Scan(&resMap)

	if r.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,Rows affected:%d", r.RowsAffected)
		return &resMap, errors.New(errMessage)
	}
	if r.Error != nil {
		return &resMap, r.Error
	}

	return &resMap, nil
}
