package requestmodels

import "mime/multipart"

type DishReq struct {
	Name                string                  `form:"name" validate:"required,gte=2"`
	CategoryId          uint                    `form:"category_id" validate:"required"`
	Description         string                  `form:"description" validate:"required,gte=5"`
	CuisineType         string                  `form:"cuisine_type" validate:"required,gte=2" json:"cuisine_type"`
	MRP                 float64                 `form:"mrp" validate:"required"`
	PortionSize         string                  `form:"portion_size"`
	DietaryInformation  string                  `form:"dietary_information"`
	Calories            int                     `form:"calories" validate:"number"`
	Protein             int                     `form:"protein" validate:"number"`
	Carbohydrates       int                     `form:"carbohydrates" validate:"number"`
	Fat                 int                     `form:"fat" validate:"number"`
	SpiceLevel          string                  `form:"spice_level" validate:"lte=20"`
	AllergenInformation string                  `form:"allergen_information" validate:"lte=50"`
	RecommendedPairings string                  `form:"recommended_pairings" validate:"lte=50"`
	SpecialFeatures     string                  `form:"special_features" validate:"lte=50"`
	Image               []*multipart.FileHeader `form:"image" validate:"required"`
	PreparationTime     string                  `form:"preparation_time" validate:"lte=15"`
	PromotionDiscount   uint                    `form:"promotion_discount"`
	StoryOrigin         string                  `form:"story_origin" validate:"lte=300"`

	Price        float64
	RestaurantId string `validate:"required"`
	ImageURL1    string
	ImageURL2    string
	ImageURL3    string
}

type DishUpdateReq struct {
	Name                string  `json:"name" validate:"required,gte=2"`
	CategoryId          uint    `json:"category_id" validate:"required"`
	Description         string  `json:"description" validate:"required,gte=5"`
	CuisineType         string  `json:"cuisine_type" validate:"required,gte=2" `
	MRP                 float64 `json:"mrp" validate:"required"`
	PortionSize         string  `json:"portion_size"`
	DietaryInformation  string  `json:"dietary_information"`
	Calories            int     `json:"calories" validate:"number"`
	Protein             int     `json:"protein" validate:"number"`
	Carbohydrates       int     `json:"carbohydrates" validate:"number"`
	Fat                 int     `json:"fat" validate:"number"`
	SpiceLevel          string  `json:"spice_level" validate:"lte=20"`
	AllergenInformation string  `json:"allergen_information" validate:"lte=50"`
	RecommendedPairings string  `json:"recommended_pairings" validate:"lte=50"`
	SpecialFeatures     string  `json:"special_features" validate:"lte=50"`
	PreparationTime     string  `json:"preparation_time" validate:"lte=15"`
	PromotionDiscount   uint    `json:"promotion_discount"`
	StoryOrigin         string  `json:"story_origin" validate:"lte=300"`
	Availability        bool    `json:"availability" validate:"required"`
	RemainingQuantity   int     `json:"quantity" validate:"required"`
	RestaurantId        string  `validate:"required"`

	Price float64
}
