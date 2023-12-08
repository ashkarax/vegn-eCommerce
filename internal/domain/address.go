package domain

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID         uint   `gorm:"not null"`
	User           Users  `gorm:"foreignKey:UserID"`
	Line1          string `gorm:"not null"`
	Street         string `gorm:"not null"`
	City           string `gorm:"not null"`
	State          string `gorm:"not null"`
	PostalCode     string `gorm:"not null"`
	Country        string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	AlternatePhone string `gorm:"not null"`
}
