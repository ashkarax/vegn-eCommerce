package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"gorm.io/gorm"
)

type CouponRepo struct {
	DB *gorm.DB
}

func NewCouponRepo(db *gorm.DB) interfaceRepository.ICouponRepository {
	return &CouponRepo{DB: db}
}

func (d *CouponRepo) GetAllCoupons() (*[]responsemodels.CouponDetails, error) {
	var couponSlice []responsemodels.CouponDetails

	query := "SELECT * FROM coupons WHERE coupon_status != 'deleted'"
	result := d.DB.Raw(query).Scan(&couponSlice)
	if result.RowsAffected == 0 {
		return &couponSlice, errors.New("no coupons present in db")
	}
	if result.Error != nil {
		return &couponSlice, result.Error
	}

	return &couponSlice, nil
}

func (d *CouponRepo) CheckCouponCodeExists(newCCode *string) error {
	var count uint
	query := "SELECT COUNT(*) FROM coupons WHERE coupon_code = ? AND coupon_status != 'deleted';"
	d.DB.Raw(query, newCCode).Scan(&count)
	if count != 0 {
		return errors.New("coupon-code is having a unique costrain,try again with another coupom-code")
	}
	return nil

}

func (d *CouponRepo) AddNewCoupon(couponData *requestmodels.CouponDetailsReq) error {
	query := "INSERT INTO coupons (coupon_code, discount_percentage, min_amount, max_amount, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?)"

	err := d.DB.Exec(query,
		couponData.CouponCode,
		couponData.DiscountPercentage,
		couponData.MinAmount,
		couponData.MaxAmount,
		couponData.StartDate,
		couponData.EndDate,
	).Error
	if err != nil {
		return err
	}
	return nil

}

func (d *CouponRepo) ChangeCouponStatus(id *string, newstat *string) error {

	query := "UPDATE coupons SET coupon_status=? WHERE coupon_id=? AND coupon_status !='deleted'"
	result := d.DB.Exec(query, newstat, id)
	if result.RowsAffected == 0 {
		fmtt := fmt.Sprintf("no coupons with id=%s present in db,it may be deleted recently", *id)
		return errors.New(fmtt)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *CouponRepo) UpdateCoupon(couponData *requestmodels.CouponDetailsReq) error {
	query := "UPDATE coupons SET coupon_code = ?, discount_percentage = ?, min_amount = ?, max_amount = ?, start_date = ?, end_date = ?  WHERE coupon_id = ? AND coupon_status !='deleted';"

	err := d.DB.Exec(query,
		couponData.CouponCode,
		couponData.DiscountPercentage,
		couponData.MinAmount,
		couponData.MaxAmount,
		couponData.StartDate,
		couponData.EndDate,
		couponData.Id,
	)
	if err.RowsAffected == 0 {
		return errors.New("enter a valid coupon-id")
	}
	if err.Error != nil {
		return err.Error
	}
	return nil

}

func (d *CouponRepo) GetcouponDetailsByCode(code *string) (*responsemodels.CouponDetails, error) {
	var couponDetails responsemodels.CouponDetails

	query := "SELECT * FROM coupons WHERE coupon_code = ? AND coupon_status = 'active';"
	result := d.DB.Raw(query, code).Scan(&couponDetails)
	if result.RowsAffected == 0 {
		return &couponDetails, errors.New("enter a valid coupon-code")
	}
	if result.Error != nil {
		return &couponDetails, result.Error
	}
	return &couponDetails, nil

}
