package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handlers.UserHandler, dish *handlers.DishHandler, address *handlers.AddressHandler, cart *handlers.CartHandler, order *handlers.OrderHandler, payment *handlers.PaymentHandler,category *handlers.CategoryHandler, JWTmiddleware *middlewares.TokenRequirements) {

	engin.POST("/signup", user.UserSignUp)
	engin.POST("/verify", user.UserOTPVerication)
	engin.POST("/login", user.UserLogin)
	// engin.POST("/forgotpassword",user.forgotpassword)
	categorymanagement := engin.Group("/category")
	{
		categorymanagement.GET("/", category.FetchActiveCategories)


	}
	paymentmanagement := engin.Group("/payment")
	{
		paymentmanagement.GET("/", payment.OnlinePayment)

	}
	dishmanagement := engin.Group("/dishes")
	{
		dishmanagement.GET("/", dish.GetAllDishesForUser)
		dishmanagement.GET("/:categoryid", dish.FetchDishesByCategoryId)
		// 	dishmanagement.GET("/search",Dish.SearchDishOrRestaurant)
	}



	engin.Use(JWTmiddleware.UserAuthorization)
	{

		addressmanagement := engin.Group("/address")
		{
			addressmanagement.GET("/", address.GetAllAddress)
			addressmanagement.PATCH("/:addressId", address.EditAddress)
			addressmanagement.POST("/", address.AddNewAddress)
		}

		profilemanagement := engin.Group("/profile")
		{
			profilemanagement.GET("/", user.GetUserProfile)
			profilemanagement.PATCH("/", user.EditUserProfile)

		}

		cartmanagement := engin.Group("/cart")
		{
			cartmanagement.POST("/:dishid/addtocart", cart.AddToCart)
			cartmanagement.GET("/", cart.GetCartDetailsOfUser)
			cartmanagement.DELETE("/:dishid", cart.DecrementorRemoveFromCart)
		}
		ordermanagement := engin.Group("/order")
		{
			ordermanagement.POST("/", order.PlaceNewOrder)
			ordermanagement.GET("/", order.GetAllOrders)
			ordermanagement.PATCH("/:ordereditemsid/cancel", order.CancelOrder)
			ordermanagement.PATCH("/:ordereditemsid/return", order.ReturnOrder)
		}
		paymentmanagement := engin.Group("/payments")
		{
			paymentmanagement.POST("/verify/:orderid", payment.VerifyPayment)

		}

	}

}
