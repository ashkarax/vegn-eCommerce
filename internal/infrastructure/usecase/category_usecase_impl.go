package usecase

import (
	"fmt"

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

func (r *CategoryUseCase) AddNewCategory(categoryData *requestmodels.CategoryReq) (*string,error) {
	var bae string
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryData)
	if err != nil {
		fmt.Println(err)
		return &bae,err
	}

	existStat := r.Repo.CheckCategoryExists(&categoryData.CategoryName)
	if existStat != nil {
		return &bae,existStat
	}

	categoryId,errr := r.Repo.AddNewCategory(categoryData)
	if errr != nil {
		return categoryId,errr
	}
	return categoryId,nil
}


func (r *CategoryUseCase)  GetAllCategories()(*[]responsemodels.CategoryRes,error){
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

func (r *CategoryUseCase)  UpdateCategory(categoryData *requestmodels.CategoryReq) (*string, error){
	var bae string
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryData)
	if err != nil {
		fmt.Println(err)
		return &bae,err
	}

	categoryId,errr := r.Repo.UpdateCategorybyId(categoryData)
	if errr != nil {
		return categoryId,errr
	}
	return categoryId,nil
}

func (r *CategoryUseCase) FetchActiveCategories()(*[]responsemodels.CategoryRes, error){
	categoriesSlice, err := r.Repo.FetchActiveCategories()
	if err != nil {
		return categoriesSlice, err
	}
	return categoriesSlice, nil
}