package routehandler

import (
	"invoice/models"
	"invoice/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetInvoiceDetails(c *gin.Context) {
	email := c.Param("email")

	details, err := service.GetInvoiceByEmail(email)
	if err != nil {
		log.Error().Err(err).Send()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoice"})
	}
	c.JSON(http.StatusOK, details)
}

func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice

	if err := c.BindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}
	if err := service.AddInvoice(&invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
	}

}
