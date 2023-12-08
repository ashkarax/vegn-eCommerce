package requestmodels

type UserSignUpReq struct {
	FirstName       string `json:"firstName" validate:"required,gte=3"`
	LastName        string `json:"lastName" validate:"required,gte=1"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,e164"`
	Password        string `json:"password" validate:"required,gte=3"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type OtpVerification struct {
	Otp string `json:"otp" validate:"required,len=4,number"`
}
type UserLoginReq struct {
	Phone    string `json:"phone"    validate:"required,gte=10,e164"`
	Password string `json:"password" validate:"required,min=4"`
}

type UserEditProf struct {
	UserId string `validate:"required"`
	FName  string `json:"firstName" validate:"required,gte=3"`
	LName  string `json:"lastName" validate:"required,gte=1"`
	Email  string `json:"email" validate:"required,email"`
}
