package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func MenuRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetAllMenus)
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenuById)
	incomingRoutes.POST("/menus", controller.CreateMenu)
	incomingRoutes.PUT("/menus/:menu_id", controller.UpdateMenu)
	incomingRoutes.DELETE("/menus/:menu_id", controller.DeleteMenu)
}
