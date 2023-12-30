package requestmodels

type SalesReportYYMMDD struct {
	Year  string `json:"year" validate:"required,lte=4"`
	Month string `json:"month" validate:"lte=2"`
	Day   string `json:"day" validate:"lte=2"`
}
