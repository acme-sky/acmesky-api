package handlers

import (
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle GET request for `Offer` model.
// It returns a list of offers.
// GetOffers godoc
//
//	@Summary	Get all offers
//	@Schemes
//	@Description	Get all offers
//	@Tags			Offers
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/offers/ [get]
func OfferHandlerGet(c *gin.Context) {
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

	var offers []models.Offer

	if user.IsAdmin {
		db.Preload("Journey").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Find(&offers)
	} else {
		db.Preload("Journey").Where("user_id = ?", userId).Omit("User").Find(&offers)
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(offers),
		"data":  &offers,
	})
}

// Handle GET request for a selected id.
// Returns the offer or a 404 status
// GetOfferById godoc
//
//	@Summary	Get an offer
//	@Schemes
//	@Description	Get an offer
//	@Tags			Offers
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/offers/{offerId}/ [get]
func OfferHandlerGetId(c *gin.Context) {
	db, _ := db.GetDb()

	var offer models.Offer
	if err := db.Where("id = ?", c.Param("id")).Preload("Journey").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
	}).First(&offer).Error; err != nil {
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

	if !(user.IsAdmin || int(user.Id) == *offer.UserId) {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, offer)
}
