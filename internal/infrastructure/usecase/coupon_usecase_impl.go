package usecase

import (
	"fmt"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/go-playground/validator/v10"
)

type CouponUseCase struct {
	Repo interfaceRepository.ICouponRepository
}

func NewCouponUseCase(repo interfaceRepository.ICouponRepository) interfaceUseCase.ICouponUseCase {
	return &CouponUseCase{Repo: repo}
}

func (r *CouponUseCase) GetAllCoupons() (*[]responsemodels.CouponDetails, error) {
	couponSlice, err := r.Repo.GetAllCoupons()
	if err != nil {
		return couponSlice, err
	}
	return couponSlice, nil
}

func (r *CouponUseCase) AddNewCoupon(couponData *requestmodels.CouponDetailsReq) error {

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(couponData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	existStat := r.Repo.CheckCouponCodeExists(&couponData.CouponCode)
	if existStat != nil {
		return existStat
	}

	errr := r.Repo.AddNewCoupon(couponData)
	if errr != nil {
		return errr
	}
	return nil

}

func (r *CouponUseCase) ChangeCouponStatus(id *string, newstat *string) error {
	errr := r.Repo.ChangeCouponStatus(id, newstat)
	if errr != nil {
		return errr
	}
	return nil
}

func (r *CouponUseCase) EditCoupon(couponData *requestmodels.CouponDetailsReq) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(couponData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	existStat := r.Repo.CheckCouponCodeExists(&couponData.CouponCode)
	if existStat != nil {
		return existStat
	}

	errr := r.Repo.UpdateCoupon(couponData)
	if errr != nil {
		return errr
	}
	return nil
}
