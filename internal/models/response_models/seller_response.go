package responsemodels

type SellerSignUpRes struct {
	RestaurantName  string `json:"restuarant_name,omitempty" `
	OwnerName       string `json:"owner_name,omitempty" `
	Email           string `json:"email,omitempty" `
	Password        string `json:"password,omitempty" `
	ConfirmPassword string `json:"confirmPassword,omitempty" `
	Description     string `json:"description,omitempty" `
	ContactNo       string `json:"contact_no,omitempty" `
	District        string `json:"district,omitempty" `
	Locality        string `json:"locality,omitempty" `
	PinCode         string `json:"pinCode,omitempty" `
	GST_NO          string `json:"gst_no,omitempty" `
	SellerExist     string `json:"sellerExist,omitempty"`
	Result          string `json:"result,omitempty"`
}

type SellerLoginRes struct {
	Phone        string `json:"phone,omitempty"`
	Password     string `json:"password,omitempty"`
	Result       string `json:"result,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}
