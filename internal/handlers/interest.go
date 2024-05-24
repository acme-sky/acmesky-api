package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/acme-sky/acmesky-api/pkg/message"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle GET request for `Interest` model.
// It returns a list of interests.
// GetInterests godoc
//
//	@Summary	Get all interests
//	@Schemes
//	@Description	Get all interests
//	@Tags			Interests
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/interests/ [get]
func InterestHandlerGet(c *gin.Context) {
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

	var interests []models.Interest

	if user.IsAdmin {
		db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
		}).Find(&interests)
	} else {
		db.Where("user_id = ?", userId).Omit("User").Find(&interests)
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(interests),
		"data":  &interests,
	})
}

// Handle POST request for `Interest` model.
// Validate JSON input by the request and crate a new interest. Finally returns
// the new created data (after preloading the foreign key objects).
// PostInterests godoc
//
//	@Summary	Create a new interest
//	@Schemes
//	@Description	Create a new interest
//	@Tags			Interests
//	@Accept			json
//	@Produce		json
//	@Success		201
//	@Router			/v1/interests/filter/ [post]
func InterestHandlerPost(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not defined"})
		return
	}

	db, _ := db.GetDb()
	var input models.InterestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.ValidateInterest(db, input, userId.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payload := map[string]interface{}{
		"flight1_departure_airport": input.Flight1DepartureAirport,
		"flight1_departure_time":    fmt.Sprintf("%sT00:00:00Z", input.Flight1DepartureTime.Format("2006-01-02")),
		"flight1_arrival_airport":   input.Flight1ArrivalAirport,
		"flight1_arrival_time":      fmt.Sprintf("%sT00:00:00Z", input.Flight1ArrivalTime.Format("2006-01-02")),
		"user_id":                   userId,
	}

	if input.Flight2DepartureTime != nil && input.Flight2DepartureAirport != nil && input.Flight2ArrivalTime != nil && input.Flight2ArrivalAirport != nil {
		payload["flight2_departure_airport"] = *input.Flight2DepartureAirport
		payload["flight2_departure_time"] = fmt.Sprintf("%sT00:00:00Z", (*input.Flight2DepartureTime).Format("2006-01-02"))
		payload["flight2_arrival_airport"] = *input.Flight2ArrivalAirport
		payload["flight2_arrival_time"] = fmt.Sprintf("%sT01:00:00Z", (*input.Flight2ArrivalTime).Format("2006-01-02"))
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":            "CM_New_Request_Save_Flight",
		"correlation_key": "0",
		"payload":         payload,
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

// Handle GET request for a selected id.
// Returns the interest or a 404 status
// GetInterestById godoc
//
//	@Summary	Get a interest
//	@Schemes
//	@Description	Get a interest
//	@Tags			Interests
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/interests/{interestId}/ [get]
func InterestHandlerGetId(c *gin.Context) {
	db, _ := db.GetDb()

	var interest models.Interest
	if err := db.Where("id = ?", c.Param("id")).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Name", "Email", "Address", "ProntogramUsername", "IsAdmin")
	}).First(&interest).Error; err != nil {
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

	if !(user.IsAdmin || int(user.Id) == interest.UserId) {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, interest)
}

// Handle DELETE request for a selected id.
// GetInterestById godoc
//
//	@Summary	Delete an interest
//	@Schemes
//	@Description	Delete a interest
//	@Tags			Interests
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Router			/v1/interests/{interestId}/ [get]
func InterestHandlerDelete(c *gin.Context) {
	db, _ := db.GetDb()

	var interest models.Interest
	if err := db.Where("id = ?", c.Param("id")).Preload("User").First(&interest).Error; err != nil {
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

	if !(user.IsAdmin || int(user.Id) == interest.UserId) {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}

	db.Delete(&interest)

	c.Status(http.StatusNoContent)
}
