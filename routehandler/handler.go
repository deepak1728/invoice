package routehandler

import (
	"context"
	"fmt"
	"invoice/models"
	"invoice/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @Summary Get invoice details by email
// @Description Retrieve invoice details for a specific user using their email.
// @Tags invoices
// @Accept json
// @Produce json
// @Param email path string true "Email of the user"
// @Success 200 {object} models.Invoice
// @Failure 500 {object} models.ErrorMsg
// @Router /invoicedetails/{email} [get]
func GetInvoiceDetails(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	filter := bson.M{}

	for key, values := range queryParams {
		if len(values) > 0 {
			filter[key] = values[0]
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	details, err := service.GetInvoiceByEmail(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
			return
		} else {
			log.Error().Err(err).Send()
			c.JSON(http.StatusInternalServerError, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
			return
		}
	}
	c.JSON(http.StatusOK, details)
}

// @Summary Create a new invoice
// @Description Create a new invoice for a specific user.
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body models.Invoice true "Invoice data"
// @Success 201 {object} models.Invoice
// @Failure 500 {object} models.ErrorMsg
// @Router /createinvoice [post]
func CreateInvoice(c *gin.Context) {
	var jsonData bson.M
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := service.AddInvoice(ctx, jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusCreated, "Successfully Created")

}
