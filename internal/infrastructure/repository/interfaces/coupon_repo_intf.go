package interfaceRepository

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type ICouponRepository interface {
	GetAllCoupons() (*[]responsemodels.CouponDetails, error)
	AddNewCoupon(*requestmodels.CouponDetailsReq) error

	CheckCouponCodeExists(*string) error
	ChangeCouponStatus(*string, *string) error
	UpdateCoupon(*requestmodels.CouponDetailsReq) error

	GetcouponDetailsByCode(*string) (*responsemodels.CouponDetails,error)

}
