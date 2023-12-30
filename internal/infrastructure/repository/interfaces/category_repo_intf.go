package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type ICategoryRepository interface {
	CheckCategoryExists(*string) error
	CheckCategoryExistsById(*string) error
	AddNewCategory(*requestmodels.CategoryReq) (*string, error)

	GetAllCategories() (*[]responsemodels.CategoryRes, error)
	ChangeCategoryStatus(*string, *string) error

	UpdateCategorybyId(*requestmodels.CategoryReq) (*string, error)

	FetchActiveCategories() (*[]responsemodels.CategoryRes, error)

	CheckCategoryOfferExists(*string, *string) error
	CreateNewCategoryOffer(*requestmodels.CategoryOfferReq) (*string, error)
	GetAllCategoryOffersByRestId(*string) (*[]responsemodels.CategoryOfferRes, error)

	UpdateCategoryOffer(*requestmodels.EditCategoryOffer) (*responsemodels.CategoryOfferRes, error)

	ChangeCategoryOfferStatus(*string, *string) error
	GetCategoryOfferStat(*string) (*string, error)
}
