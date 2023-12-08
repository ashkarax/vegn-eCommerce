package responsemodels

type CategoryRes struct {
	CategoryID     string `json:"category_id,omitempty"`
	CategoryName   string `json:"category_name,omitempty"`
	CategoryStatus string `json:"category_status,omitempty"`
}
