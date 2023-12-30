package usecase

import (
	"errors"
	"fmt"
	"time"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/go-playground/validator/v10"
)

type CategoryUseCase struct {
	Repo interfaceRepository.ICategoryRepository
}

func NewCategoryUseCase(repo interfaceRepository.ICategoryRepository) interfaceUseCase.ICategoryUseCase {
	return &CategoryUseCase{Repo: repo}
}

func (r *CategoryUseCase) AddNewCategory(categoryData *requestmodels.CategoryReq) (*string, error) {
	var bae string //just to match the return type of this function,when only error is returned
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryData)
	if err != nil {
		fmt.Println(err)
		return &bae, err
	}

	existStat := r.Repo.CheckCategoryExists(&categoryData.CategoryName)
	if existStat != nil {
		return &bae, existStat
	}

	categoryId, errr := r.Repo.AddNewCategory(categoryData)
	if errr != nil {
		return categoryId, errr
	}
	return categoryId, nil
}

func (r *CategoryUseCase) GetAllCategories() (*[]responsemodels.CategoryRes, error) {
	categoriesSlice, err := r.Repo.GetAllCategories()
	if err != nil {
		return categoriesSlice, err
	}
	return categoriesSlice, nil
}

func (r *CategoryUseCase) ChangeCategoryStatus(id *string, newstat *string) error {
	errr := r.Repo.ChangeCategoryStatus(id, newstat)
	if errr != nil {
		return errr
	}
	return nil
}

func (r *CategoryUseCase) UpdateCategory(categoryData *requestmodels.CategoryReq) (*string, error) {
	var bae string
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryData)
	if err != nil {
		fmt.Println(err)
		return &bae, err
	}

	categoryId, errr := r.Repo.UpdateCategorybyId(categoryData)
	if errr != nil {
		return categoryId, errr
	}
	return categoryId, nil
}

func (r *CategoryUseCase) FetchActiveCategories() (*[]responsemodels.CategoryRes, error) {
	categoriesSlice, err := r.Repo.FetchActiveCategories()
	if err != nil {
		return categoriesSlice, err
	}
	return categoriesSlice, nil
}

//restaurant side

func (r *CategoryUseCase) AddNewCategoryOffer(categoryOfferData *requestmodels.CategoryOfferReq) (*string, error) {
	var fillerString string //just to match the return type of this function,when only error is returned
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryOfferData)
	if err != nil {
		fmt.Println(err)
		return &fillerString, err
	}

	existStat0 := r.Repo.CheckCategoryExistsById(&categoryOfferData.CategoryID)
	if existStat0 != nil {
		return &fillerString, existStat0
	}

	existStat := r.Repo.CheckCategoryOfferExists(&categoryOfferData.CategoryID, &categoryOfferData.RestaurantID)
	if existStat != nil {
		return &fillerString, existStat
	}

	EndDate := time.Now().Add(time.Duration(categoryOfferData.Validity) * 24 * time.Hour)
	categoryOfferData.EndDate = EndDate

	categoryOfferId, errr := r.Repo.CreateNewCategoryOffer(categoryOfferData)
	if errr != nil {
		return &fillerString, errr
	}
	return categoryOfferId, nil
}

func (r *CategoryUseCase) GetAllCategoryOffers(restId *string) (*[]responsemodels.CategoryOfferRes, error) {
	categoryOffersSlice, err := r.Repo.GetAllCategoryOffersByRestId(restId)
	if err != nil {
		return categoryOffersSlice, err
	}
	return categoryOffersSlice, nil
}

func (r *CategoryUseCase) EditCategoryOffer(categoryOfferData *requestmodels.EditCategoryOffer) (*responsemodels.CategoryOfferRes, error) {
	var fillerStruct responsemodels.CategoryOfferRes
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryOfferData)
	if err != nil {
		fmt.Println(err)
		return &fillerStruct, err
	}

	EndDate := time.Now().Add(time.Duration(categoryOfferData.Validity) * 24 * time.Hour)
	categoryOfferData.EndDate = EndDate

	updatedRes, errr := r.Repo.UpdateCategoryOffer(categoryOfferData)
	if errr != nil {
		return updatedRes, errr
	}
	return updatedRes, nil
}

func (r *CategoryUseCase) ChangeCategoryOfferStatus(categoryOfferData *requestmodels.CategoryOfferStatus) error {

	if categoryOfferData.Status == "block" {
		currentStat, err := r.Repo.GetCategoryOfferStat(&categoryOfferData.CategoryOfferId)
		if err != nil {
			return err
		}
		if *currentStat != "active" {
			return errors.New("can only block a category-offer which is active")
		}
		newStatus := "blocked"
		errUpdt := r.Repo.ChangeCategoryOfferStatus(&categoryOfferData.CategoryOfferId, &newStatus)
		if errUpdt != nil {
			return errUpdt
		}

	}

	if categoryOfferData.Status == "unblock" {
		currentStat, err := r.Repo.GetCategoryOfferStat(&categoryOfferData.CategoryOfferId)
		if err != nil {
			return err
		}
		if *currentStat != "blocked" {
			return errors.New("can only unblock a category-offer which is blocked")
		}
		newStatus := "active"
		errUpdt := r.Repo.ChangeCategoryOfferStatus(&categoryOfferData.CategoryOfferId, &newStatus)
		if errUpdt != nil {
			return errUpdt
		}

	}

	if categoryOfferData.Status == "delete" {
		newStatus := "deleted"
		errUpdt := r.Repo.ChangeCategoryOfferStatus(&categoryOfferData.CategoryOfferId, &newStatus)
		if errUpdt != nil {
			return errUpdt
		}


	}

	return nil

}
