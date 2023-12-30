package domain

import "time"

type Category struct {
	CategoryID     uint `gorm:"primarykey"`
	CategoryName   string
	CategoryStatus status `gorm:"default:active"`
}

type CategoryOffer struct {
	CategoryOfferID    uint `gorm:"primarykey"`
	OfferTitle         string
	CategoryId         uint       `gorm:"not null"`
	Category           Category   `gorm:"foreignKey:CategoryId"`
	RestaurantID       uint       `gorm:"not null"`
	Restaurant         Restaurant `gorm:"foreignKey:RestaurantID"`
	DiscountPercentage float64
	StartDate          time.Time
	EndDate            time.Time
	OfferStatus        status `gorm:"default:active"`
}
