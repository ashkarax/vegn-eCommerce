package responsemodels

import "time"

type CartItemInfo struct {
	CartID             uint      `json:"cart_id" `
	Name               string    `json:"productname" `
	DishID             string    `json:"DishID"`
	RestaurantID       string    `json:"RestaurantID"`
	RestaurantName     string    `json:"restaurant_name"`
	Quantity           uint      `json:"quantity"`
	MRP                float64   `json:"MRP"`
	PromotionDiscount  uint      `json:"promotion_discount"`
	Price              float64   `json:"price_after_promotion_discount"`
	OfferTitle         string    `json:"category_discount_title,omitempty"`
	DiscountPercentage uint      `json:"discountPercentage,omitempty"`
	SalePrice          float64   `json:"sale_price"`
	EndDate            time.Time `json:"end_date,omitempty"`
	ImageURL1          string    `json:"image_url1" `
	Availability       bool      `json:"availability" `

	RemainingQuantity uint `json:"null,omitempty"`
	CategoryOfferID   uint `json:"null,omitempty"`
}

type CartDetailsResp struct {
	UserID      string          `json:"user_id" `
	TotalPrice  float64         `json:"total_price"`
	DishesCount uint            `json:"Dish_count"`
	Cart        *[]CartItemInfo `json:"cart"`
}
