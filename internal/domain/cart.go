package domain

type Cart struct {
	CartID      uint       `gorm:"primarykey"`
	UserID       uint       `gorm:"not null"`
	Users        Users      `gorm:"foreignKey:UserID"`
	RestaurantID uint       `gorm:"not null"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
	DishID       uint       `gorm:"not null"`
	Dish         Dish       `gorm:"foreignKey:DishID"`
	Quantity     int        `gorm:"not null"`
}
