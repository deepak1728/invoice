package routehandler

import (
	"fmt"
	"invoice/models"
	"invoice/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
	email := c.Param("email")

	details, err := service.GetInvoiceByEmail(email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
		}
		log.Error().Err(err).Send()
		c.JSON(http.StatusInternalServerError, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
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
	var invoice models.Invoice

	if err := c.BindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
		return
	}
	if err := service.AddInvoice(&invoice); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMsg{Message: fmt.Sprintf("%v", err)})
		return
	}

}
