package handlers

import (
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle GET request for `AvailableFlight` model.
// It returns a list of availableFlights.
// GetAvailableFlights godoc
//
//	@Summary	Get all availableFlights
//	@Schemes
//	@Description	Get all available flights
//	@Tags			AvailableFlights
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/available-flights/ [get]
func AvailableFlightHandlerGet(c *gin.Context) {
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

	var availableFlights []models.AvailableFlight

	if user.IsAdmin {
		db.Preload("Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Find(&availableFlights)
	} else {
		db.Preload("Interest").Where("user_id = ?", userId).Omit("User").Find(&availableFlights)
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(availableFlights),
		"data":  &availableFlights,
	})
}

// Handle GET request for a selected id.
// Returns the available flight or a 404 status
// GetAvailableFlightById godoc
//
//	@Summary	Get an availble flight
//	@Schemes
//	@Description	Get an available flight
//	@Tags			AvailableFlights
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/available-flights/{availableFlightId}/ [get]
func AvailableFlightHandlerGetId(c *gin.Context) {
	db, _ := db.GetDb()

	var availableFlight models.AvailableFlight
	if err := db.Where("id = ?", c.Param("id")).Preload("Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
	}).First(&availableFlight).Error; err != nil {
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

	if !(user.IsAdmin || int(user.Id) == *availableFlight.UserId) {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, availableFlight)
}
