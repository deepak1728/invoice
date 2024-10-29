package routes

import (
	"invoice/routehandler"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/invoicedetails/:email", routehandler.GetInvoiceDetails)
	router.POST("/createinvoice", routehandler.CreateInvoice)

	return router
}
