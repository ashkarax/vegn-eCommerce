package domain

import "gorm.io/gorm"

type status string

const (
	Blocked status = "blocked"
	Deleted status = "deleted"
	Pending status = "pending"

	Active status = "active"

	//only for restaurants
	verified status = "verified"
	Rejected status = "rejected"
)

type Users struct {
	gorm.Model
	FName    string
	LName    string
	Email    string
	Phone    string
	Password string
	Status   status  `gorm:"default:pending"`
	Wallet   float64 `gorm:"default:0"`
}
