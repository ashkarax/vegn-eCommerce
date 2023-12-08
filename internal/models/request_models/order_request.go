package requestmodels

type OrderDetails struct {
	UserID        string `validate:"required"`
	AddressID     string `json:"address_id" validate:"required,gte=1,lte=5"`
	PaymentMethod string `json:"Payment_method" validate:"required,lte=20"`
	CouponCode    string `json:"coupon_code" validate:"lte=20"`

	OrderStatus     string
	PaymentStatus   string
	OrderIdRazorPay string
	TotalAmount     float64
	CouponId        uint
}

type CanOrRetReq struct {
	UserID         string
	OrderedItemsID string
	RestaurantID   string
}
