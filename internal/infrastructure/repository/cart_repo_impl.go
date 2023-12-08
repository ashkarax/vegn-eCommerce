package repository

import (
	"errors"
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type CartRepo struct {
	DB *gorm.DB
}

func NewCartRepo(DB *gorm.DB) interfaceRepository.ICartRepository {
	return &CartRepo{DB: DB}
}

func (d *CartRepo) AddToCart(userIdString *string, dishId *string, restaurantId *string) error {
	query := "INSERT INTO carts (user_id, restaurant_id, dish_id, quantity) VALUES (?,?,?,?)"
	result := d.DB.Exec(query, userIdString, restaurantId, dishId, 1)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *CartRepo) CheckDishAlreadyInCart(userIdString *string, dishId *string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM carts WHERE user_id = ? AND dish_id = ?"
	err := d.DB.Raw(query, userIdString, dishId).Scan(&count)
	if err.Error != nil {
		return false, err.Error
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (d *CartRepo) IncrementDishCountInCart(userIdString *string, dishId *string) error {
	query := "UPDATE carts SET quantity = quantity + 1 WHERE user_id = ? AND dish_id = ?"
	err := d.DB.Exec(query, userIdString, dishId)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (d *CartRepo) FetchCartItemsofUser(userID *string) (*[]responsemodels.CartItemInfo, error) {
	var cartItemsInfoSlice []responsemodels.CartItemInfo

	query := "SELECT carts.cart_id,carts.quantity,carts.dish_id,dishes.name,dishes.restaurant_id,dishes.price,dishes.image_url1,dishes.image_url2,dishes.image_url3,dishes.availability,dishes.remaining_quantity FROM carts INNER JOIN dishes ON carts.dish_id=dishes.id WHERE carts.user_id=?"
	result := d.DB.Raw(query, userID).Scan(&cartItemsInfoSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user have no cart,cart empty")
	}
	return &cartItemsInfoSlice, nil

}

func (d *CartRepo) DeleteFromCart(dishId *string, userid *string) error {
	query := "DELETE FROM carts WHERE dish_id =? AND user_id=?"
	result := d.DB.Exec(query, dishId, userid)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user have no cart")
	}
	return nil
}

func (d *CartRepo) ReturnQuantityOfCartItem(userIdString *string, dishId *string) (*int, error) {
	var quantity int
	query := "SELECT quantity FROM carts WHERE user_id = ? AND dish_id = ?"
	err := d.DB.Raw(query, userIdString, dishId).Scan(&quantity)
	if err.Error != nil {
		return &quantity, err.Error
	}

	return &quantity, nil

}
func (d *CartRepo) DecrementDishCountInCart(userIdString *string, dishId *string) error {
	query := "UPDATE carts SET quantity = quantity - 1 WHERE user_id = ? AND dish_id = ?"
	err := d.DB.Exec(query, userIdString, dishId)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (d *CartRepo) DeleteCartAfterOrder(CartID *uint) error {
	query := "DELETE FROM carts WHERE cart_id=?"
	err := d.DB.Exec(query, CartID)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
