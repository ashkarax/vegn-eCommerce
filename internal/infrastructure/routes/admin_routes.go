package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"

	// "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handlers.AdminHandler, coupon *handlers.CouponHandler, category *handlers.CategoryHandler, JWTmiddleware *middlewares.TokenRequirements) {
	engin.POST("/login", admin.AdminLogin)

	engin.Use(JWTmiddleware.AdminAuthorization)
	{

		restaurantmanagement := engin.Group("/restaurants")
		{

			restaurantmanagement.GET("/", admin.VerifiedRestuarants)
			restaurantmanagement.PATCH("/block/:id", admin.BlockRestaurant)

			restaurantmanagement.GET("/pending", admin.PendingRestuarants)
			restaurantmanagement.PATCH("/pending/verify/:id", admin.VerifyRestaurant)
			restaurantmanagement.PATCH("/pending/reject/:id", admin.RejectRestaurant)

			restaurantmanagement.GET("/blocked", admin.BlockedRestuarants)
			restaurantmanagement.PATCH("/blocked/unblock/:id", admin.UnBlockRestaurant)
			restaurantmanagement.PATCH("/blocked/delete/:id", admin.DeleteRestaurant)

			restaurantmanagement.GET("/rejected", admin.RejectedRestuarants)
		}
		usermanagement := engin.Group("/users")
		{

			usermanagement.GET("/", admin.LatestUsers)
			usermanagement.GET("/search", admin.SearchUser)

			usermanagement.PATCH("/block/:id", admin.BlockUser)

			usermanagement.GET("/blocked", admin.BlockedUsers)
			usermanagement.PATCH("/blocked/unblock/:id", admin.UnBlockUser)

		}
		couponmanagement := engin.Group("/coupon")
		{
			couponmanagement.GET("/", coupon.AllCoupons)
			couponmanagement.POST("/", coupon.AddNewCoupon)
			couponmanagement.DELETE("/:couponid", coupon.DeleteCoupon)
			couponmanagement.PATCH("/block/:couponid", coupon.BlockCoupon)
			couponmanagement.PATCH("/unblock/:couponid", coupon.UnBlockCoupon)
			couponmanagement.PATCH("/edit/:couponid", coupon.EditCoupon)

		}
		categorymanagement := engin.Group("/category")
		{
			categorymanagement.POST("/", category.NewCategory)
			categorymanagement.GET("/", category.FetchAllCategory)
			categorymanagement.PATCH("/:categoryid", category.UpdateCategory)
			categorymanagement.DELETE("/:categoryid", category.DeleteCategory)

		}

	}

}
