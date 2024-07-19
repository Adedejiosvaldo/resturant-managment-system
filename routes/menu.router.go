package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/menu", controllers.CreateAMenu)
	router.GET("/menu/:menu_id", controllers.GetAMenu)
	router.POST("/menu", controllers.CreateAMenu)
	router.PATCH("/menu/:menu_id", controllers.UpdateAMenu)
	router.DELETE("/menu/:menu_id", controllers.DeleteAMenu)
}
