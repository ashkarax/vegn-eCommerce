package responsemodels

type AddressRes struct {
	Id             uint   `json:"address_id,omitempty"`
	UserID         uint   `json:"user_id,omitempty"`
	Line1          string `json:"line1,omitempty"`
	Street         string `json:"street,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	PostalCode     string `json:"postal_code,omitempty"`
	Country        string `json:"country,omitempty"`
	Phone          string `json:"phone,omitempty"`
	AlternatePhone string `json:"alternate_phone,omitempty"`
}
