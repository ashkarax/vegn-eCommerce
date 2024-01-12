package routes

import (
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/handlers"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(engin *gin.RouterGroup, Restaurant *handlers.RestaurantHandler, Dish *handlers.DishHandler, order *handlers.OrderHandler, category *handlers.CategoryHandler, JWTmiddleware *middlewares.TokenRequirements) {

	engin.POST("/signup", Restaurant.RestaurantSignUp)            //swagprogress
	engin.POST("/login", Restaurant.RestaurantLogin)                //swagprogress

	engin.Use(JWTmiddleware.RestaurantAuthorization)
	{
		dishmanagement := engin.Group("/dish")
		{

			dishmanagement.GET("/", Dish.FetchAllDishesForRestaurant)          //swagprogress
			dishmanagement.POST("/", Dish.NewDish)                              //swagprogress //need modifications

			dishmanagement.GET("/:dishid", Dish.FetchDishWithId)
			dishmanagement.PATCH("/:dishid", Dish.UpdateDishDetails)            //swagprogress
			dishmanagement.DELETE("/:dishid", Dish.DeleteDish)
		}
		ordermanagement := engin.Group("/orders")
		{
			ordermanagement.GET("/", order.FetchAllOrdersForRestaurant)                 //swagprogress
			// ordermanagement.GET("/:orderid",order.FetchOrderWithId)
			ordermanagement.PATCH("/status", order.ChangeStatus)                          //swagprogress
		}

		salesreportmanagement:=engin.Group("/salesreport")
		{ 
			salesreportmanagement.POST("/",order.GetSalesreporYYMMDD)                            //swagprogress
			salesreportmanagement.GET("/:customdays",order.GetSalesreporCustomDays)               //swagprogress
			salesreportmanagement.GET("/xlsx",order.GenerateSalesReportXlsx)                       //swagprogress
		}
		categorymanagement := engin.Group("/category")
		{
			categorymanagement.GET("/", category.FetchActiveCategories)                             //swagprogress 
			categorymanagement.GET("/offer", category.GetAllCategoryOffer)                          //swagprogress 
			categorymanagement.POST("/offer", category.CreateCategoryOffer)                         //swagprogress 
            categorymanagement.PATCH("/offer/:categoryofferid",category.EditCategoryOffer)          //swagprogress 
			categorymanagement.PATCH("/offer",category.ChangeCategoryOfferStatus)                   //swagprogress 

		}
	}
}
