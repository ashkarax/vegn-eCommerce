package requestmodels

type RestaurantSignUpReq struct {
	RestaurantName  string `json:"restaurant_name" validate:"required,gte=2"`
	OwnerName       string `json:"owner_name" validate:"required,gte=2"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,gte=4"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
	Description     string `json:"description" validate:"required,gte=10"`
	Phone           string `json:"phone" validate:"required,e164"`
	District        string `json:"district" validate:"required,gte=2"`
	Locality        string `json:"locality" validate:"required,gte=2"`
	PinCode         string `json:"pinCode" validate:"required,len=6"`
	GST_NO          string `json:"gst_no" validate:"required,gte=2"`
}
type RestaurantLoginReq struct {
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=4"`
}


