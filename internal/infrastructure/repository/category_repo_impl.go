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
func (d *CategoryRepo) CheckCategoryExistsById(categoryId *string) error {
	var count uint
	query := "SELECT COUNT(*) FROM categories WHERE category_id = ? AND category_status = 'active';"
	d.DB.Raw(query, categoryId).Scan(&count)
	if count == 0 {
		return errors.New("category does not exist or not active")
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

// restaurant side
func (d *CategoryRepo) CheckCategoryOfferExists(CategoryID *string, RestaurantID *string) error {
	var count uint
	query := "SELECT COUNT(*) FROM category_offers WHERE category_id = ? AND restaurant_id=? AND offer_status = 'active' AND end_date> now();"
	d.DB.Raw(query, CategoryID, RestaurantID).Scan(&count)
	if count != 0 {
		return errors.New("restaurant already have an existing offer on this category,try editing it or delete the old one")
	}
	return nil

}

func (d *CategoryRepo) CreateNewCategoryOffer(categ *requestmodels.CategoryOfferReq) (*string, error) {
	var categoryOfferId string
	query := "INSERT INTO category_offers (offer_title,category_id,restaurant_id,discount_percentage,start_date,end_date) VALUES (?,?,?,?,NOW(),?) RETURNING category_offer_id;"

	err := d.DB.Raw(query,
		categ.Title, categ.CategoryID, categ.RestaurantID, categ.CategoryDiscount, categ.EndDate,
	).Scan(&categoryOfferId).Error
	if err != nil {
		return &categoryOfferId, err
	}

	return &categoryOfferId, nil
}

func (d *CategoryRepo) GetAllCategoryOffersByRestId(restId *string) (*[]responsemodels.CategoryOfferRes, error) {
	var categoryOffersSlice []responsemodels.CategoryOfferRes

	query := "SELECT * FROM category_offers WHERE restaurant_id=? AND offer_status='active' AND end_date> now()"
	result := d.DB.Raw(query, restId).Scan(&categoryOffersSlice)
	if result.RowsAffected == 0 {
		return &categoryOffersSlice, errors.New("no active offers exists for this seller")
	}
	if result.Error != nil {
		return &categoryOffersSlice, result.Error
	}

	return &categoryOffersSlice, nil

}

func (d *CategoryRepo) UpdateCategoryOffer(categ *requestmodels.EditCategoryOffer) (*responsemodels.CategoryOfferRes, error) {
	var updatedRes responsemodels.CategoryOfferRes
	query := "UPDATE category_offers SET offer_title=?,discount_percentage=?,start_date=NOW(),end_date=? RETURNING *;"

	err := d.DB.Raw(query,
		categ.Title, categ.CategoryDiscount, categ.EndDate,
	).Scan(&updatedRes).Error
	if err != nil {
		return &updatedRes, err
	}

	return &updatedRes, nil
}

func (d *CategoryRepo) ChangeCategoryOfferStatus(id *string, newStat *string) error {
	query := "UPDATE category_offers SET offer_status=? WHERE category_offer_id=?"
	err := d.DB.Exec(query, newStat, id).Error
	if err != nil {
		return err
	}
	return nil

}

func (d *CategoryRepo) GetCategoryOfferStat(id *string) (*string, error) {
	var stat string
	query := "SELECT offer_status FROM category_offers WHERE category_offer_id=? AND offer_status!='deleted' AND end_date >now()"
	res := d.DB.Raw(query, id).Scan(&stat)
	if res.Error != nil {
		return &stat, res.Error
	}
	if res.RowsAffected == 0 {
		return &stat, errors.New("Category-Offer with this id does not exist,already deleted or expired")
	}
	return &stat, nil
}
