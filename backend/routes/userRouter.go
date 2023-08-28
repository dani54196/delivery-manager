package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetAllUsers)
	incomingRoutes.GET("/users/:user_id", controller.GetUserById)
	incomingRoutes.POST("/users/signup", controller.SignUp)
	incomingRoutes.POST("/users/login", controller.Login)
}
