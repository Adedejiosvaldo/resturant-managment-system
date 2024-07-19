package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	router.GET("/tables", controllers.GetTables)
	router.GET("/tables/:table_id", controllers.GetATable)
	router.POST("/tables", controllers.CreateATable)
	router.PATCH("/tables/:table_id", controllers.UpdateATable)
	router.DELETE("/tables/:table_id", controllers.DeleteATable)
}
