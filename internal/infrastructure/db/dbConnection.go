package db

import (
	"fmt"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	"github.com/ashkarax/vegn-eCommerce/internal/domain"
	hashpassword "github.com/ashkarax/vegn-eCommerce/pkg/hash_password"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config config.DataBase) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dberr != nil {
		return DB, nil
	}

	// Table Creation
	if err := DB.AutoMigrate(&domain.Admin{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Users{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Restaurant{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Dish{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Address{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Cart{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Order{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.OrderedItems{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Coupon{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Category{}); err != nil {
		return DB, err
	}
	CheckAndCreateAdmin(DB)
	return DB, nil

}
func CheckAndCreateAdmin(DB *gorm.DB) {
	var count int
	var (
		Name     = "Vegn"
		Email    = "vegn@gmail.com"
		Password = "vegnvegn"
	)
	HashedPassword := hashpassword.HashPassword(Password)

	query := "SELECT COUNT(*) FROM admins"
	DB.Raw(query).Row().Scan(&count)
	if count <= 0 {
		query = "INSERT INTO admins(name, email, password) VALUES(?, ?, ?)"
		DB.Exec(query, Name, Email, HashedPassword).Row().Err()
	}
}
