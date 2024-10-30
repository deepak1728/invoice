package routes

import (
	"invoice/routehandler"

	_ "invoice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title  Documenting API
// @version 1

// @localhost:8181

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/invoicedetails/:email", routehandler.GetInvoiceDetails)
	router.POST("/createinvoice", routehandler.CreateInvoice)

	return router
}
