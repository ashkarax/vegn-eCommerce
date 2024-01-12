package requestmodels

type AddressReq struct {
	UserID    string 
	AddressId string 

	Line1          string `json:"line1" validate:"required"`
	Street         string `json:"street" validate:"required"`
	City           string `json:"city" validate:"required"`
	State          string `json:"state" validate:"required"`
	PostalCode     string `json:"postal_code" validate:"required"`
	Country        string `json:"country" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	AlternatePhone string `json:"alternate_phone" validate:"required"`
}
