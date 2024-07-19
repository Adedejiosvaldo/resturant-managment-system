package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func NoteRoute(router *gin.Engine) {
	router.GET("/notes", controllers.GetNotes)
	router.GET("/notes/:note_id", controllers.GetANote)
	router.POST("/notes", controllers.CreateANote)
	router.PATCH("/notes/:note_id", controllers.UpdateANote)
	router.DELETE("/note/:note_id", controllers.DeleteANote)

}
