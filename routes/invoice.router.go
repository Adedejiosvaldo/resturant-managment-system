package routes

import (
	"github.com/adedeji/resturant-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(router *gin.Engine) {
	router.GET("/invoices", controllers.GetInvoices)
	router.GET("/invoices/:invoice_id", controllers.GetInvoice)
	router.POST("/invoices", controllers.CreateInvoice)
	router.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice)
	router.DELETE("/invoices/:invoice_id", controllers.DeleteInvoice)
}
