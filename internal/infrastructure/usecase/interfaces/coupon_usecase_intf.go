package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
)

type ICouponUseCase interface {
	GetAllCoupons() (*[]responsemodels.CouponDetails,error)
	AddNewCoupon(*requestmodels.CouponDetailsReq) error

	ChangeCouponStatus(*string,*string) error
	EditCoupon(*requestmodels.CouponDetailsReq) error
	
}
