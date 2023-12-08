package responsemodels

type CartItemInfo struct {
	CartID            uint    `json:"cart_id" `
	Name              string  `json:"productname" `
	DishID            string  `json:"DishID"`
	RestaurantID      string  `json:"RestaurantID"`
	Quantity          uint    `json:"quantity"`
	Price             float64 `json:"price"`
	ImageURL1         string  `json:"image_url1" `
	ImageURL2         string  `json:"image_url2,omitempty" `
	ImageURL3         string  `json:"image_url3,omitempty" `
	Availability      bool    `json:"availability" `
	RemainingQuantity uint    `json:"remaining_quantity,omitempty" `
}

type CartDetailsResp struct {
	UserID      string          `json:"user_id" `
	TotalPrice  float64         `json:"total_price"`
	DishesCount uint            `json:"Dish_count"`
	Cart        *[]CartItemInfo `json:"cart"`
}
