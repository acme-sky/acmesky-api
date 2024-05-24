package handlers

import (
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle GET request for `Journey` model.
// It returns a list of journeys.
// GetJourneys godoc
//
//	@Summary	Get all journeys
//	@Schemes
//	@Description	Get all journeys
//	@Tags			Journeys
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/journeys/ [get]
func JourneyHandlerGet(c *gin.Context) {
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

	var journeys []models.Journey

	if user.IsAdmin {
		db.Preload("Flight1").Preload("Flight2").Preload("Flight1.Interest").Preload("Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Find(&journeys)
	} else {
		db.Preload("Flight1").Preload("Flight2").Preload("Flight1.Interest").Preload("Flight2.Interest").Where("user_id = ?", userId).Omit("User").Find(&journeys)
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(journeys),
		"data":  &journeys,
	})
}

// Handle GET request for a selected id.
// Returns the journey or a 404 status
// GetJourneyById godoc
//
//	@Summary	Get an journey
//	@Schemes
//	@Description	Get an journey
//	@Tags			Journeys
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/journeys/{journeyId}/ [get]
func JourneyHandlerGetId(c *gin.Context) {
	db, _ := db.GetDb()

	var journey models.Journey
	if err := db.Where("id = ?", c.Param("id")).Preload("Flight1").Preload("Flight2").Preload("Flight1.Interest").Preload("Flight2.Interest").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
	}).First(&journey).Error; err != nil {
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

	if !(user.IsAdmin || int(user.Id) == *journey.UserId) {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, journey)
}
