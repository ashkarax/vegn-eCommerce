package domain

import "gorm.io/gorm"

type status string

const (
	Active  status = "active"
	Blocked status = "blocked"
	Deleted status = "deleted"
	Pending status = "pending"
)

type Users struct {
	gorm.Model
	FName    string
	LName    string
	Email    string
	Phone    string
	Password string
	Status   status `gorm:"default:pending"`
}
