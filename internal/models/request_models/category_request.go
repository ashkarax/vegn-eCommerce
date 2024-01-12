package requestmodels

import "time"

type CategoryReq struct {
	CategoryName string `json:"category_name" validate:"required"`

	CategoryId string
}

type CategoryOfferReq struct {
	Title            string `json:"title" validate:"required,lte=100"`
	CategoryID       string `json:"category_id" validate:"required,number"`
	CategoryDiscount uint   `json:"category_discount" validate:"required,min=1,max=99"`
	Validity         uint   `json:"validityindays" validate:"required,min=0"`

	RestaurantID string
	EndDate      time.Time
}

type EditCategoryOffer struct {
	Title            string `json:"title" validate:"required"`
	CategoryDiscount uint   `json:"category_discount" validate:"required,min=1,max=99"`
	Validity         uint   `json:"validityindays"`

	CategoryOfferID string
	RestaurantID    string
	EndDate         time.Time
}

type CategoryOfferStatus struct {
	Status          string `json:"status" validate:"required,lte=20"`
	CategoryOfferId string `json:"id" validate:"required"`
}
