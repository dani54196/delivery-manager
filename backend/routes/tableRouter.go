package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func TableRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetAllTables)
	incomingRoutes.GET("/tables/:table_id", controller.GetTableById)
	incomingRoutes.POST("/tables", controller.CreateTable)
	incomingRoutes.PUT("/tables/:table_id", controller.UpdateTable)
	incomingRoutes.DELETE("/tables/:table_id", controller.DeleteTable)
}
