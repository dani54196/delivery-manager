package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRouter(router)
	router.Use(middlewares.Authtentication())

	router.FoodRouter(router)
	router.MenuRouter(router)
	router.TableRouter(router)
	router.OrderRouter(router)
	router.OrderItemRouter(router)
	router.InvoiceRouter(router)

	router.Run(":" + port)

}
