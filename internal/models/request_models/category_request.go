package requestmodels

type CategoryReq struct {
	CategoryName string `json:"category_name" validate:"required"`

	CategoryId string
}
