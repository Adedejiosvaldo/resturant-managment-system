package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:order_id", controllers.GetAOrder)
	router.POST("/orders", controllers.CreateAOrder)
	router.PATCH("/orders/:order_id", controllers.UpdateAOrder)
	router.DELETE("/orders/:order_id", controllers.DeleteAOrder)
}
