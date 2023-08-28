package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func NoteRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controller.GetAllNotes)
	incomingRoutes.GET("/notes/:note_id", controller.GetNoteById)
	incomingRoutes.POST("/notes", controller.CreateNote)
	incomingRoutes.PUT("/notes/:note_id", controller.UpdateNote)
	incomingRoutes.DELETE("/notes/:note_id", controller.DeleteNote)
}
