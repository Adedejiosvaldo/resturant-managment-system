package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:user_id", controllers.GetUsers)
	router.POST("/users/signup", controllers.SignUp)
	router.POST("/user/login", controllers.Login)
}
