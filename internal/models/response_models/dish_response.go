package responsemodels

type DishRes struct {
	ID                  uint    `json:"id,omitempty"`
	RestaurantID        uint    `json:"restaurant_id,omitempty"`
	CategoryId          uint    `json:"category_id,omitempty"`
	Name                string  `json:"name,omitempty"`
	Description         string  `json:"description,omitempty"`
	CuisineType         string  `json:"cuisine_type,omitempty"`
	MRP                 float64 `json:"mrp,omitempty"`
	PortionSize         string  `json:"portion_size,omitempty"`
	DietaryInformation  string  `json:"dietary_information,omitempty"`
	Calories            int     `json:"calories,omitempty"`
	Protein             int     `json:"protein,omitempty"`
	Carbohydrates       int     `json:"carbohydrates,omitempty"`
	Fat                 int     `json:"fat,omitempty"`
	SpiceLevel          string  `json:"spice_level,omitempty"`
	AllergenInformation string  `json:"allergen_information,omitempty"`
	RecommendedPairings string  `json:"recommended_pairings,omitempty"`
	SpecialFeatures     string  `json:"special_features,omitempty"`
	ImageURL1           string  `json:"image_url1,omitempty" `
	ImageURL2           string  `json:"image_url2,omitempty" `
	ImageURL3           string  `json:"image_url3,omitempty" `
	PreparationTime     string  `json:"preparation_time,omitempty"`
	PromotionDiscount   string  `json:"promotion_discount,omitempty"`
	Price               float64 `json:"price,omitempty"`
	StoryOrigin         string  `json:"story_origin,omitempty"`
	Availability        bool    `json:"availability,omitempty"`
	RemainingQuantity   int     `json:"quantity,omitempty"`

	CategoryOfferID    string `json:"category_offer_id,omitempty"`
	OfferTitle         string `json:"title,omitempty"`
	DiscountPercentage uint   `json:"discountPercentage,omitempty"`

	SalePrice float64 `json:"sale_price,omitempty"`
}
