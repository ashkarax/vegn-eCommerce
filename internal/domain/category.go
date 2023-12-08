package domain

type Category struct {
	CategoryID     uint `gorm:"primarykey"`
	CategoryName   string
	CategoryStatus status `gorm:"default:active"`
}
