package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(engin *gin.RouterGroup, Restaurant *handlers.RestaurantHandler, Dish *handlers.DishHandler, order *handlers.OrderHandler, category *handlers.CategoryHandler, JWTmiddleware *middlewares.TokenRequirements) {

	engin.POST("/signup", Restaurant.RestaurantSignUp)
	engin.POST("/login", Restaurant.RestaurantLogin)

	engin.Use(JWTmiddleware.RestaurantAuthorization)
	{
		dishmanagement := engin.Group("/dish")
		{

			dishmanagement.GET("/", Dish.FetchAllDishesForRestaurant)
			dishmanagement.POST("/", Dish.NewDish)

			dishmanagement.GET("/:dishid", Dish.FetchDishWithId)
			dishmanagement.PATCH("/:dishid", Dish.UpdateDishDetails)
			dishmanagement.DELETE("/:dishid", Dish.DeleteDish)
		}
		ordermanagement := engin.Group("/orders")
		{
			ordermanagement.GET("/", order.FetchAllOrdersForRestaurant)
			// ordermanagement.GET("/:orderid",order.FetchOrderWithId)
			ordermanagement.PATCH("/:ordereditemsid/preparing",order.ChangeStatusToPreparing)
			ordermanagement.PATCH("/:ordereditemsid/outfordeliver",order.ChangeStatusToOutForDelivery)
			ordermanagement.PATCH("/:ordereditemsid/deliver",order.ChangeStatusToDelivered)
			// ordermanagement.PATCH("/:ordereditemsid/cancel",order.RestaurantCancelOrder)

		}
		categorymanagement := engin.Group("/category")
		{
			categorymanagement.GET("/", category.FetchActiveCategories)

		}
	}
}
