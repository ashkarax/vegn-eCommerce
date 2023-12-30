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
			ordermanagement.PATCH("/status", order.ChangeStatus)
		}

		salesreportmanagement:=engin.Group("/salesreport")
		{
			salesreportmanagement.GET("/",order.GetSalesreporYYMMDD)
			salesreportmanagement.GET("/:customdays",order.GetSalesreporCustomDays)
			salesreportmanagement.GET("/xlsx",order.GenerateSalesReportXlsx)
		}
		categorymanagement := engin.Group("/category")
		{
			categorymanagement.GET("/", category.FetchActiveCategories)

			categorymanagement.GET("/offer", category.GetAllCategoryOffer)
			categorymanagement.POST("/offer", category.CreateCategoryOffer)
            categorymanagement.PATCH("/offer/:categoryofferid",category.EditCategoryOffer)
			categorymanagement.PATCH("/offer",category.ChangeCategoryOfferStatus)

		}
	}
}
