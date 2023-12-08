package responsemodels

import "time"

type OrderDetailsRes struct {
	AddressID     string `json:"address_id,omitempty"`
	PaymentMethod string `json:"Payment_method,omitempty"`
	CouponCode    string `json:"coupon_code,omitempty" `

	OrderId          string  `json:"order_id,omitempty" `
	TotalAmount      float64 `json:"total_anount,omitempty" `
	DiscountedAmount float64 `json:"amount_after_applying_coupon,omitempty" `
}

type OrderDetailsResponse struct {
	OrderedItemsID uint    `json:"ordered_items_id,omitempty"`
	OrderID        uint    `json:"order_id,omitempty"`
	DishID         uint    `json:"DishID,omitempty"`
	Name           string  `json:"Dish_name,omitempty" `
	OrderQuantity  uint    `json:"quantity,omitempty"`
	DishPrice      float64 `json:"price,omitempty"`
	RestaurantID   uint    `json:"RestaurantID,omitempty"`
	RestaurantName string  `json:"Restaurant_name,omitempty" `
	PaymentMethod  string  `json:"Payment_method,omitempty"`
	ImageURL1      string  `json:"ImageURl1,omitempty" `

	AddressID      uint      `json:"address_id,omitempty"`
	Line1          string    `json:"Address_line1,omitempty" `
	Phone          string    `json:"contact_no,omitempty" `
	Street         string    `json:"Street,omitempty" `
	City           string    `json:"City,omitempty" `
	OrderStatus    string    `json:"order_status,omitempty"`
	OrderDate      time.Time `json:"order_date,omitempty"`
	DeliverDate    time.Time `json:"delivery_date,omitempty"`
	PostalCode     string    `json:"pincode,omitempty" `
	AlternatePhone string    `json:"AlternatePhone,omitempty" `
	TotalPrice     float64   `json:"total_price,omitempty"`
}

type CanOrRetResp struct {
	PaymentMethod string
	OrderStatus   string
	PaymentStatus string

	OrderQuantity uint
	DishPrice     float64
	DishID        uint
	RestaurantId  uint
	DeliverDate   time.Time
}
