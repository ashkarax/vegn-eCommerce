package responsemodels

import "time"

type CategoryRes struct {
	CategoryID     string `json:"category_id,omitempty"`
	CategoryName   string `json:"category_name,omitempty"`
	CategoryStatus string `json:"category_status,omitempty"`
}

type CategoryOfferRes struct {
	CategoryOfferID    string    `json:"category_offer_id"`
	OfferTitle         string    `json:"title,omitempty"`
	CategoryId         uint      `json:"category_id,omitempty"`
	RestaurantID       uint      `json:"restaurant_id,omitempty"`
	DiscountPercentage uint      `json:"discountPercentage,omitempty"`
	StartDate          time.Time `json:"start_date,omitempty"`
	EndDate            time.Time `json:"end_date,omitempty"`
	StatusOfferStatus  string    `json:"status,omitempty"`
}
