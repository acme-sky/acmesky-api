package handlers

import (
	"net/http"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"
)

// Handle GET request for `User` model.
// It returns a list of user.
// GetUser godoc
//
//	@Summary	Get all users
//	@Schemes
//	@Description	Get all users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/users/ [get]
func UserHandlerGet(c *gin.Context) {
	db, _ := db.GetDb()

	var users []models.User
	db.Omit("Password").Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"data":  &users,
	})
}

// Handle GET request for a selected id.
// Returns the user or a 404 status
// GetUserById godoc
//
//	@Summary	Get an user
//	@Schemes
//	@Description	Get an user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/users/{userId}/ [get]
func UserHandlerGetId(c *gin.Context) {
	db, _ := db.GetDb()
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).Omit("Password").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Handle PUT request for `User` model.
// This endpoint should be called only by the owner or by `is_admin=1` user.
// EditUserById godoc
//
//	@Summary	Edit an user
//	@Schemes
//	@Description	Edit an user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/users/{userId}/ [put]
func UserHandlerPut(c *gin.Context) {
	db, _ := db.GetDb()
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Omit("Password").Model(&user).Updates(input)

	c.JSON(http.StatusOK, user)
}
