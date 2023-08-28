package routes

import (
	"restaurant-manager/controller"

	"github.com/gin-gonic/gin"
)

func InvoiceRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetAllInvoices)
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoiceById)
	incomingRoutes.POST("/invoices", controller.CreateInvoice)
	incomingRoutes.PUT("/invoices/:invoice_id", controller.UpdateInvoice)
	incomingRoutes.DELETE("/invoices/:invoice_id", controller.DeleteInvoice)
}
