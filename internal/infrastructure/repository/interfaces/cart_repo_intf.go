package interfaceRepository

import responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"

type ICartRepository interface{
	AddToCart(*string,*string,*string) error
	IncrementDishCountInCart(*string,*string) error
	CheckDishAlreadyInCart(*string,*string) (bool,error)
	FetchCartItemsofUser(*string)(*[]responsemodels.CartItemInfo,error)
	ReturnQuantityOfCartItem(*string,*string) (*int,error) 
	DecrementDishCountInCart(*string,*string) error
	DeleteFromCart(*string,*string) error

	DeleteCartAfterOrder(*uint) error

}