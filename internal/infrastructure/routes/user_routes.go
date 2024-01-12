package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handlers.UserHandler, dish *handlers.DishHandler, address *handlers.AddressHandler, cart *handlers.CartHandler, order *handlers.OrderHandler, payment *handlers.PaymentHandler, category *handlers.CategoryHandler, JWTmiddleware *middlewares.TokenRequirements) {

	engin.POST("/signup", user.UserSignUp)                                                       //swagprogress
	engin.POST("/verify", user.UserOTPVerication)                                                //swagprogress
	engin.POST("/login", user.UserLogin)                                                       //swagprogress
	// engin.POST("/forgotpassword",user.forgotpassword)

	categorymanagement := engin.Group("/category")
	{ 
		categorymanagement.GET("/", category.FetchActiveCategories)                               //swagprogress

	}
	paymentmanagement := engin.Group("/payment")
	{
		paymentmanagement.GET("/", payment.OnlinePayment)

	}
	dishmanagement := engin.Group("/dishes")
	{
		dishmanagement.GET("/", dish.GetAllDishesForUser)                                    //swagprogress
		dishmanagement.GET("/:categoryid", dish.FetchDishesByCategoryId)                       //swagprogress
		// 	dishmanagement.GET("/search",Dish.SearchDishOrRestaurant)
	}

	engin.Use(JWTmiddleware.UserAuthorization)
	{

		addressmanagement := engin.Group("/address")
		{
			addressmanagement.GET("/", address.GetAllAddress)                                 //swagprogress
			addressmanagement.PATCH("/:addressId", address.EditAddress)                        //swagprogress
			addressmanagement.POST("/", address.AddNewAddress)                                  //swagprogress
		}

		profilemanagement := engin.Group("/profile")
		{ 
			profilemanagement.GET("/", user.GetUserProfile)                                         //swagprogress
			profilemanagement.PATCH("/", user.EditUserProfile)                                      //swagprogress

		}

		cartmanagement := engin.Group("/cart")
		{ 
			cartmanagement.POST("/addtocart/:dishid", cart.AddToCart)                                //swagprogress
			cartmanagement.GET("/", cart.GetCartDetailsOfUser)                                        //swagprogress
			cartmanagement.DELETE("/:dishid", cart.DecrementorRemoveFromCart)                           //swagprogress
		}
		ordermanagement := engin.Group("/order")
		{
			ordermanagement.POST("/", order.PlaceNewOrder)                                               //swagprogress           
			ordermanagement.GET("/", order.GetAllOrders)                                                 //swagprogress
			ordermanagement.PATCH("/cancel/:ordereditemsid", order.CancelOrder)
			ordermanagement.PATCH("/return/:ordereditemsid", order.ReturnOrder)
			ordermanagement.GET("/invoice/:orderid", order.GenerateInvoice)                               //swagprogress 

		}
		paymentmanagement := engin.Group("/payments")
		{
			paymentmanagement.POST("/verify/:orderid", payment.VerifyPayment)                              //swagprogress 

		}

	}

}
