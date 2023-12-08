package domain

import "time"

type Coupon struct {
	CouponID           uint `gorm:"primarykey"`
	CouponCode         string
	DiscountPercentage float64
	MinAmount          float64
	MaxAmount          float64
	StartDate          time.Time
	EndDate            time.Time
	CouponStatus       status `gorm:"default:active"`
}
