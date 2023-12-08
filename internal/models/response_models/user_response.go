package responsemodels

type SignupData struct {
	FName           string `json:"first_name,omitempty"`
	LName           string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Password        string `json:"password,omitempty"`
	OTP             string `json:"otp,omitempty"`
	Token           string `json:"token,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	IsUserExist     string `json:"isUserExist,omitempty"`
}

type OtpVerifResult struct {
	Phone        string `json:"phone,omitempty"`
	Otp          string `json:"otp,omitempty"`
	Result       string `json:"result,omitempty"`
	Token        string `json:"token,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}

type UserLoginRes struct {
	Phone        string `json:"phone,omitempty"`
	Password     string `json:"password,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}

type UserDetails struct {
	Id        int    `json:"user_id,omitempty"`
	FName     string `json:"first_name,omitempty"`
	LName     string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
}
