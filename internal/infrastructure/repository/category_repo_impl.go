package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) interfaceRepository.ICategoryRepository {
	return &CategoryRepo{DB: db}
}

func (d *CategoryRepo) CheckCategoryExists(categoryName *string) error {
	var count uint
	query := "SELECT COUNT(*) FROM categories WHERE category_name = ? AND category_status != 'deleted';"
	d.DB.Raw(query, categoryName).Scan(&count)
	if count != 0 {
		return errors.New("category-name is having a unique costrain,try again with another category-name")
	}
	return nil

}

func (d *CategoryRepo) AddNewCategory(categ *requestmodels.CategoryReq) (*string, error) {
	var categoryId string

	query := "INSERT INTO categories (category_name) VALUES (?) RETURNING category_id;"

	err := d.DB.Raw(query,
		categ.CategoryName,
	).Scan(&categoryId).Error
	if err != nil {
		return &categoryId, err
	}

	return &categoryId, nil
}

func (d *CategoryRepo) GetAllCategories() (*[]responsemodels.CategoryRes, error) {
	var categorySlice []responsemodels.CategoryRes

	query := "SELECT * FROM categories WHERE category_status!='deleted'"
	result := d.DB.Raw(query).Scan(&categorySlice)
	if result.RowsAffected == 0 {
		return &categorySlice, errors.New("no categories present in db")
	}
	if result.Error != nil {
		return &categorySlice, result.Error
	}

	return &categorySlice, nil

}

func (d *CategoryRepo) ChangeCategoryStatus(id *string, newstat *string) error {

	query := "UPDATE categories SET category_status=? WHERE category_id=? AND category_status !='deleted'"
	result := d.DB.Exec(query, newstat, id)
	if result.RowsAffected == 0 {
		fmtt := fmt.Sprintf("no category with id=%s present in db,it may be deleted recently", *id)
		return errors.New(fmtt)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *CategoryRepo) UpdateCategorybyId(categ *requestmodels.CategoryReq) (*string, error) {
	categoryId := categ.CategoryId

	query := "UPDATE categories SET category_name=? WHERE category_id=? AND category_status!='deleted'"

	err := d.DB.Exec(query, categ.CategoryName, categ.CategoryId)
	if err.RowsAffected == 0 {
		fmtt := fmt.Sprintf("no category with id=%s present in db,it may be deleted recently", categ.CategoryId)
		return &categoryId, errors.New(fmtt)
	}
	if err.Error != nil {
		return &categoryId, err.Error
	}

	return &categoryId, nil
}

func (d *CategoryRepo) FetchActiveCategories() (*[]responsemodels.CategoryRes, error) {
	var categorySlice []responsemodels.CategoryRes

	query := "SELECT * FROM categories WHERE category_status='active'"
	result := d.DB.Raw(query).Scan(&categorySlice)
	if result.RowsAffected == 0 {
		return &categorySlice, errors.New("no categories present in db")
	}
	if result.Error != nil {
		return &categorySlice, result.Error
	}

	return &categorySlice, nil

}
