package usecase

import (
	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type CartUseCase struct {
	CartRepo interfaceRepository.ICartRepository
	DishRepo interfaceRepository.IDishRepo
}

func NewCartUseCase(repo interfaceRepository.ICartRepository, dishRepository interfaceRepository.IDishRepo) interfaceUseCase.ICartUseCase {
	return &CartUseCase{CartRepo: repo, DishRepo: dishRepository}
}

func (r *CartUseCase) AddToCart(userIdString *string, dishId *string) error {

	alreadyExistStat, erro := r.CartRepo.CheckDishAlreadyInCart(userIdString, dishId)
	if erro != nil {
		return erro
	}
	if alreadyExistStat {
		errIncr := r.CartRepo.IncrementDishCountInCart(userIdString, dishId)
		if errIncr != nil {
			return errIncr
		}
		return nil
	}

	restaurantId, err := r.DishRepo.ReturnRestaurantIdofDish(dishId)
	if err != nil {
		return err
	}

	errAdd := r.CartRepo.AddToCart(userIdString, dishId, restaurantId)
	if errAdd != nil {
		return errAdd
	}
	return nil

}

func (r *CartUseCase) GetCartDetails(userIdString *string) (*responsemodels.CartDetailsResp, error) {
	var cartResp responsemodels.CartDetailsResp

	cartItemsInfo, err := r.CartRepo.FetchCartItemsofUser(userIdString)
	if err != nil {
		return &cartResp, err
	}

	totalPrice := 0.0
	var quantitySum uint 
	for _, item := range *cartItemsInfo {
		totalPrice += float64(item.Quantity) * item.Price
		quantitySum+=item.Quantity
	}

	cartResp.Cart = cartItemsInfo
	cartResp.UserID = *userIdString
	cartResp.DishesCount = quantitySum
	cartResp.TotalPrice = totalPrice
	return &cartResp, nil
}

func (r *CartUseCase) DeleteFromCart(dishId *string,userId *string) error{

	quantity, erro := r.CartRepo.ReturnQuantityOfCartItem(userId, dishId)
	if erro != nil {
		return erro
	}

	if *quantity>1 {
		errIncr := r.CartRepo.DecrementDishCountInCart(userId, dishId)
		if errIncr != nil {
			return errIncr
		}
		return nil
	}

	err := r.CartRepo.DeleteFromCart(dishId,userId)
	if err != nil {
		return err
	}
	return nil
}
