package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type ICategoryUseCase interface {
	AddNewCategory(*requestmodels.CategoryReq) (*string, error)
	GetAllCategories() (*[]responsemodels.CategoryRes, error)

	ChangeCategoryStatus(*string, *string) error
	UpdateCategory(*requestmodels.CategoryReq) (*string, error)
	FetchActiveCategories() (*[]responsemodels.CategoryRes, error)

	AddNewCategoryOffer(*requestmodels.CategoryOfferReq) (*string, error)
	GetAllCategoryOffers(*string) (*[]responsemodels.CategoryOfferRes, error)

	EditCategoryOffer(*requestmodels.EditCategoryOffer) (*responsemodels.CategoryOfferRes, error)
	ChangeCategoryOfferStatus(*requestmodels.CategoryOfferStatus) error
}
