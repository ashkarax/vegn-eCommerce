package responsemodels

import "time"

type CouponDetails struct {
	CouponID           uint      `json:"coupon_id,omitempty"`
	CouponCode         string    `json:"coupon_code,omitempty"`
	DiscountPercentage float64   `json:"discount_percentage,omitempty"`
	MinAmount          float64   `json:"min_amount,omitempty"`
	MaxAmount          float64   `json:"max_amount,omitempty"`
	StartDate          time.Time `json:"start_date,omitempty"`
	EndDate            time.Time `json:"end_date,omitempty"`
	CouponStatus       string    `json:"coupon_status,omitempty"`
}
