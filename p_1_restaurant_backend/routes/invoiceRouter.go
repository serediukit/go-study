package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant_backend/controllers"
)

func InvoiceRoutes(router *gin.Engine) {
	router.GET("/invoices", controllers.GetInvoices())
	router.GET("/foodss/:invoice_id", controllers.GetInvoice())
	router.POST("/invoices", controllers.CreateInvoice())
	router.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
