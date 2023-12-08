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
			restaurantmanagement.PATCH("/:id/block", admin.BlockRestaurant)

			restaurantmanagement.GET("/pending", admin.PendingRestuarants)
			restaurantmanagement.PATCH("/pending/:id/verify", admin.VerifyRestaurant)
			restaurantmanagement.PATCH("/pending/:id/reject", admin.RejectRestaurant)

			restaurantmanagement.GET("/blocked", admin.BlockedRestuarants)
			restaurantmanagement.PATCH("/blocked/:id/unblock", admin.UnBlockRestaurant)
			restaurantmanagement.PATCH("/blocked/:id/delete", admin.DeleteRestaurant)

			restaurantmanagement.GET("/rejected", admin.RejectedRestuarants)
		}
		usermanagement := engin.Group("/users")
		{

			usermanagement.GET("/", admin.LatestUsers)
			usermanagement.GET("/search", admin.SearchUser)

			usermanagement.PATCH("/:id/block", admin.BlockUser)

			usermanagement.GET("/blocked", admin.BlockedUsers)
			usermanagement.PATCH("/blocked/:id/unblock", admin.UnBlockUser)

		}
		couponmanagement := engin.Group("/coupon")
		{
			couponmanagement.GET("/", coupon.AllCoupons)
			couponmanagement.POST("/", coupon.AddNewCoupon)
			couponmanagement.DELETE("/:couponid", coupon.DeleteCoupon)
			couponmanagement.PATCH("/:couponid/block", coupon.BlockCoupon)
			couponmanagement.PATCH("/:couponid/unblock", coupon.UnBlockCoupon)
			couponmanagement.PATCH("/:couponid/edit", coupon.EditCoupon)

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
