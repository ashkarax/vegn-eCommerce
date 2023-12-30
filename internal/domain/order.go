package domain

import (
	"time"
)

type orderstatus string
type paymentmethod string

// const (
// 	Pending        orderstatus = "pending"
// 	Processing     orderstatus = "processing"
// 	Preparing      orderstatus = "preparing"
// 	OutForDelivery orderstatus = "outfordelivery"
// 	Delivered      orderstatus = "delivered"
// 	Cancelled      orderstatus = "cancelled"
// 	Return         orderstatus = "return"
// )

// const (
// 	COD    paymentmethod = "COD"
// 	ONLINE paymentmethod = "ONLINE"
// 	ONLINE paymentmethod = "WALLET"
// )

type Order struct {
	OrderID uint `gorm:"primarykey"`

	UserID uint  `gorm:"not null"`
	Users  Users `gorm:"foreignKey:UserID"`

	AddressID uint    `gorm:"not null"`
	Address   Address `gorm:"foreignKey:AddressID"`

	CouponId uint `gorm:"default:0"`
	// Coupn    Coupon `gorm:"foreignkey:CouponId;association_foreignkey:CouponID"`

	PaymentMethod paymentmethod

	RazorPayId string `gorm:"default:nil"`

	OrderDate time.Time `gorm:"not null"`
}

type OrderedItems struct {
	OrderedItemsID uint `gorm:"primarykey"`

	OrderID uint  `gorm:"not null"`
	Order   Order `gorm:"foreignKey:OrderID"`

	DishID uint `gorm:"not null"`
	Dish   Dish `gorm:"foreignKey:DishID"`

	OrderQuantity uint

	MRP               float64 `gorm:"default:0"`
	PromotionDiscount uint    `gorm:"default:0"`

	CategoryOfferId    uint    `gorm:" default:0"`
	DiscountPercentage float64 `gorm:" default:0"`

	DishPrice float64

	RestaurantID uint       `gorm:"not null"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`

	OrderStatus   orderstatus `gorm:"default:pending"`
	PaymentStatus string

	DeliverDate time.Time
}
