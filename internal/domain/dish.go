package domain

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	RestaurantID        uint       `gorm:"not null"`
	Restaurant          Restaurant `gorm:"foreignKey:RestaurantID"`
	CategoryId          uint
	Category            Category `gorm:"foreignKey:CategoryId"`
	Name                string   `gorm:"not null"`
	Description         string   `gorm:"type:varchar(150)"`
	CuisineType         string   `gorm:"type:varchar(50)"`
	Price               float64  `gorm:"not null"`
	PortionSize         string   `gorm:"type:varchar(50)" `
	DietaryInformation  string   `gorm:"type:varchar(100)" `
	Calories            int
	Protein             int
	Carbohydrates       int
	Fat                 int
	SpiceLevel          string `gorm:"type:varchar(20)" `
	AllergenInformation string `gorm:"type:varchar(255)" `
	RecommendedPairings string `gorm:"type:varchar(100)" `
	SpecialFeatures     string `gorm:"type:varchar(50)"`
	ImageURL1           string `gorm:"type:varchar(555)" `
	ImageURL2           string `gorm:"type:varchar(555)" `
	ImageURL3           string `gorm:"type:varchar(555)" `
	PreparationTime     string `gorm:"type:varchar(50)"`
	PromotionDiscount   string `gorm:"type:varchar(80)"`
	StoryOrigin         string `gorm:"type:varchar(300)"`
	Availability        bool   `gorm:"default:false"`
	RemainingQuantity   int    `gorm:"default:0"`
}
