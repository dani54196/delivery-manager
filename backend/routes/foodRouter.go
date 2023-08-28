package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func FoodRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetAllFoods)
	incomingRoutes.GET("/foods/:food_id", controller.GetFoodById)
	incomingRoutes.POST("/foods", controller.CreateFood)
	incomingRoutes.PUT("/foods/:food_id", controller.UpdateFood)
	incomingRoutes.DELETE("/foods/:food_id", controller.DeleteFood)
}
