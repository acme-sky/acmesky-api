package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/acme-sky/acmesky-api/pkg/message"
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
		db.Where("is_used = 't'").Preload("Journey").Preload("Journey.Flight1").Preload("Journey.Flight2").Preload("Journey.Flight1.Interest").Preload("Journey.Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Find(&offers)
	} else {
		db.Where("is_used = 't'").Preload("Journey").Preload("Journey.Flight1").Preload("Journey.Flight2").Preload("Journey.Flight1.Interest").Preload("Journey.Flight2.Interest").Where("user_id = ?", userId).Omit("User").Find(&offers)
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
	if err := db.Where("id = ?", c.Param("id")).Preload("Journey").Preload("Journey.Flight1").Preload("Journey.Flight2").Preload("Journey.Flight1.Interest").Preload("Journey.Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
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

// Confirm a received offer by the code.
//
//	@Summary	Create a new interest
//	@Schemes
//	@Description	Confirm an offer
//	@Tags			Offer
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/offers/confirm/ [post]
func OfferConfirmHandlerPost(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not defined"})
		return
	}

	db, _ := db.GetDb()

	var input models.OfferConfirmInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.ValidateOfferToken(db, input, userId.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var offer models.Offer
	if err := db.Where("token = ?", input.Token).First(&offer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":            "CM_Check_Offer",
		"correlation_key": "0",
		"payload": map[string]string{
			"token": input.Token,
		},
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := message.SendMessage(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Status(200)
}

// Handle an offer payment. It ignores which payment comes from because the
// BPMN resolves everything.
// PayOfferById godoc
//
//	@Summary	Pay an offer
//	@Schemes
//	@Description	Pay an offer
//	@Tags			Offer
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/offers/{offerId}/pay/ [post]
func OfferHandlerPay(c *gin.Context) {
	db, _ := db.GetDb()

	var offer models.Offer
	if err := db.Where("id = ? AND payment_paid = 'f'", c.Param("id")).First(&offer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":            "CM_Payment_Response",
		"correlation_key": "0",
		"payload": map[string]string{
			"payment_status": "OK",
		},
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := message.SendMessage(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Status(200)
}
