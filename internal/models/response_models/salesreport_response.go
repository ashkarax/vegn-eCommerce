package responsemodels

type SalesReportData struct {
	Orders   uint `json:"total_orders"`
	Quantity uint `json:"total_units_sold"`
	Price    uint `json:"total_revenue"`
}
