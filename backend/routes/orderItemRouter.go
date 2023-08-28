package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetAllOrderItems)
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItemById)
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem)
	incomingRoutes.PUT("/orderItems/:orderItem_id", controller.UpdateOrderItem)
	incomingRoutes.DELETE("/orderItems/:orderItem_id", controller.DeleteOrderItem)
}
