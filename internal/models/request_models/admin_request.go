package requestmodels

type AdminLoginData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4"`
}
