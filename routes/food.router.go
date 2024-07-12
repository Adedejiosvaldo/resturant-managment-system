package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/foods", controllers.GetFoods())
	router.GET("/foods/:food_id", controllers.GetAFood())
	router.POST("/foods", controllers.CreateAFood())
	router.PATCH("/foods/:food_id", controllers.UpdateAMeal())
	router.DELETE("/food/:food_id", controllers.DeleteAMeal())
}
