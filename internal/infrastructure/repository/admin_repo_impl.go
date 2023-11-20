package repository

import (
	"errors"

	interfaceRepository "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaceRepository.IAdminRepository {
	return &adminRepository{DB: db}

}

func (d *adminRepository) GetPassword(email string) (string, error) {

	var hashedPassword string

	query := "SELECT password FROM admins WHERE email= ?"
	err := d.DB.Raw(query, email).Row().Scan(&hashedPassword)
	if err != nil {
		return "", errors.New("error at admin password fetch")
	}
	return hashedPassword, nil
}
