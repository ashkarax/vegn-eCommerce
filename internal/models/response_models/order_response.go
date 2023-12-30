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

type OrderDetailsPDF struct {
	OrderDate time.Time

	FName string
	LName string
	Email string
	Phone string

	CouponCode         string
	DiscountPercentage float64

	Line1          string
	Street         string
	City           string
	State          string
	PostalCode     string
	Country        string
	AlternatePhone string
}

type OrderedItemsDataPDF struct {
	Name               string
	OrderQuantity      uint
	Restaurant_name    string
	MRP                float64
	PromotionDiscount  uint
	DiscountPercentage float64
	SalePrice          float64
	TotalAmount        float64
}

type OrderDetaisForSalesReport struct {
	OrderID            uint
	DishID             uint
	Name               string
	OrderQuantity      uint
	MRP                float64
	PromotionDiscount  uint
	DiscountPercentage float64
	DishPrice          float64
	PaymentMethod      string
	OrderDate          time.Time
	DeliverDate        time.Time
}

type OrderResponseX struct {
	OrderID        uint                `json:"order_id"`
	TotalAmount    float64             `json:"total_amount"`
	PaymentMethod  string              `json:"Payment_method"`
	PaymentStatus  string              `json:"payment_status"`
	OrderDate      time.Time           `json:"order_date"`
	FName          string              `json:"first_name"`
	LName          string              `json:"last_name"`
	Line1          string              `json:"line1"`
	Street         string              `json:"street"`
	City           string              `json:"city"`
	PostalCode     string              `json:"postal_code"`
	Phone          string              `json:"phone"`
	AlternatePhone string              `json:"alternate_phone,omitempty"`
	OrderedItems   []OrderDishDetailsX `json:"ordered_items"`
}

type OrderDishDetailsX struct {
	OrderID        uint    `json:"order_id"`
	OrderedItemsID uint    `json:"ordered_items_id"`
	DishID         uint    `json:"dish_id"`
	Name           string  `json:"dish_name"`
	ImageURL1      string  `json:"ImageURl1,omitempty" `
	PortionSize    string  `json:"portion_size,omitempty"`
	OrderQuantity  uint    `json:"order_quantity"`
	DishPrice      float64 `json:"dish_price"`
	OrderStatus    string  `json:"order_status"`
}

type OrderDataFetcherX struct {
	OrderID        uint
	PaymentMethod  string
	PaymentStatus  string
	OrderDate      time.Time
	FName          string
	LName          string
	Line1          string
	Street         string
	City           string
	State          string
	PostalCode     string
	Country        string
	Phone          string
	AlternatePhone string
	OrderedItemsID uint
	DishID         uint
	Name           string
	ImageURL1      string
	PortionSize    string
	OrderQuantity  uint
	DishPrice      float64
	OrderStatus    string
}
