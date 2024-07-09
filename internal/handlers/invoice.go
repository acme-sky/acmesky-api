package handlers

import (
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle GET request for `Invoice` model.
// It returns a list of invoices.
// GetInvoices godoc
//
//	@Summary	Get all invoices
//	@Schemes
//	@Description	Get all invoices
//	@Tags			Invoices
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/invoices/ [get]
func InvoiceHandlerGet(c *gin.Context) {
	db, _ := db.GetDb()
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not defined"})
		return
	}

	var user models.User
	if err := db.Where("id = ?", userId).Omit("Password").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	var invoices []models.Invoice

	if user.IsAdmin {
		db.Preload("Journey").Preload("Journey.Flight1").Preload("Journey.Flight2").Preload("Journey.Flight1.Interest").Preload("Journey.Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Find(&invoices)
	} else {
		db.Preload("Journey").Preload("Journey.Flight1").Preload("Journey.Flight2").Preload("Journey.Flight1.Interest").Preload("Journey.Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Where("user_id = ?", userId).Find(&invoices)
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(invoices),
		"data":  &invoices,
	})
}

// Handle GET request for a selected id.
// Returns the invoice or a 404 status
// GetInvoiceById godoc
//
//	@Summary	Get an invoice
//	@Schemes
//	@Description	Get an invoice
//	@Tags			Invoices
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/invoices/{invoiceId}/ [get]
func InvoiceHandlerGetId(c *gin.Context) {
	db, _ := db.GetDb()

	var invoice models.Invoice
	if err := db.Where("id = ?", c.Param("id")).Preload("Journey").Preload("Journey.Flight1").Preload("Journey.Flight2").Preload("Journey.Flight1.Interest").Preload("Journey.Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
	}).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	var user *models.User
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not defined"})
		return
	}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	if !(user.IsAdmin || int(user.Id) == invoice.UserId) {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, invoice)
}
