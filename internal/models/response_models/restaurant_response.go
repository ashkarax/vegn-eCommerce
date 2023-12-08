package responsemodels



type RestaurantSignUpRes struct {
	RestaurantName  string `json:"restuarant_name,omitempty" `
	OwnerName       string `json:"owner_name,omitempty" `
	Email           string `json:"email,omitempty" `
	Password        string `json:"password,omitempty" `
	ConfirmPassword string `json:"confirmPassword,omitempty" `
	Description     string `json:"description,omitempty" `
	Phone           string `json:"phone,omitempty" `
	District        string `json:"district,omitempty" `
	Locality        string `json:"locality,omitempty" `
	PinCode         string `json:"pinCode,omitempty" `
	GST_NO          string `json:"gst_no,omitempty" `
	RestaurantExist string `json:"restaurantExist,omitempty"`
	Result          string `json:"result,omitempty"`
}

type RestaurantLoginRes struct {
	Phone        string `json:"phone,omitempty"`
	Password     string `json:"password,omitempty"`
	Result       string `json:"result,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}

type RestuarntDetails struct {
	Id              int    `json:"restaurant_id,omitempty"`
	Restaurant_name string `json:"restaurant_name,omitempty"`
	Owner_name      string `json:"owner_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Description     string `json:"description,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Country         string `json:"country,omitempty"`
	State           string `json:"state,omitempty"`
	District        string `json:"district,omitempty"`
	Locality        string `json:"locality,omitempty"`
	GST_NO          string `json:"gst_no,omitempty"`
	PinCode         string `json:"pinCode,omitempty"`
	Status          string `json:"status,omitempty"`
	CreatedAt       int   `json:"created_at,omitempty"`
}
