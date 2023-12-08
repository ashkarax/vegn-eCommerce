package requestmodels

import "time"

type CouponDetailsReq struct {
	CouponCode         string    `json:"coupon_code" validate:"required"`
	DiscountPercentage float64   `json:"discount_percentage" validate:"required"`
	MinAmount          float64   `json:"min_amount" validate:"required"`
	MaxAmount          float64   `json:"max_amount" validate:"required"`
	StartDate          time.Time `json:"start_date" validate:"required"`
	EndDate            time.Time `json:"end_date" validate:"required"`

	Id string
}
