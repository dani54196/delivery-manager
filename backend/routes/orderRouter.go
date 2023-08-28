package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func OrderRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetAllOrders)
	incomingRoutes.GET("/orders/:order_id", controller.GetOrderById)
	incomingRoutes.POST("/orders", controller.CreateOrder)
	incomingRoutes.PUT("/orders/:order_id", controller.UpdateOrder)
	incomingRoutes.DELETE("/orders/:order_id", controller.DeleteOrder)
}
