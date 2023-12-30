package domain

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Restaurant_name string
	Owner_name      string
	Email           string
	Password        string
	Description     string
	Phone           string `gorm:"unique"`
	Country         string `gorm:"default:india"`
	State           string `gorm:"default:kerala"`
	District        string
	Locality        string
	GST_NO          string `gorm:"not null"`
	PinCode         string
	Status          status  `gorm:"default:pending"`
	CodWallet       float64 `gorm:"default:0.00"`
	AdminCredit     float64 `gorm:"default:0.00"` //OnlineWallet
}
