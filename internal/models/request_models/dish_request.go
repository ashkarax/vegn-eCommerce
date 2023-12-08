package requestmodels

import "mime/multipart"

type DishReq struct {
	Name                string                  `form:"name" validate:"required,gte=2"`
	CategoryId          uint                    `form:"category_id" validate:"required"`
	Description         string                  `form:"description" validate:"required,gte=5"`
	CuisineType         string                  `form:"cuisine_type" validate:"required,gte=2" json:"cuisine_type"`
	Price               float64                 `form:"price" validate:"required,gte=2,number"`
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
	PromotionDiscount   string                  `form:"promotion_discount" validate:"lte=50"`
	StoryOrigin         string                  `form:"story_origin" validate:"lte=300"`

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
	Price               float64 `json:"price" validate:"required,gte=2,number"`
	PortionSize         string  `form:"portion_size"`
	DietaryInformation  string  `form:"dietary_information"`
	Calories            int     `form:"calories" validate:"number"`
	Protein             int     `form:"protein" validate:"number"`
	Carbohydrates       int     `form:"carbohydrates" validate:"number"`
	Fat                 int     `form:"fat" validate:"number"`
	SpiceLevel          string  `form:"spice_level" validate:"lte=20"`
	AllergenInformation string  `form:"allergen_information" validate:"lte=50"`
	RecommendedPairings string  `form:"recommended_pairings" validate:"lte=50"`
	SpecialFeatures     string  `form:"special_features" validate:"lte=50"`
	PreparationTime     string  `form:"preparation_time" validate:"lte=15"`
	PromotionDiscount   string  `form:"promotion_discount" validate:"lte=50"`
	StoryOrigin         string  `form:"story_origin" validate:"lte=300"`
	Availability        bool    `json:"availability" validate:"required"`
	RemainingQuantity   int     `json:"quantity" validate:"required"`
	RestaurantId        string  `validate:"required"`
}
