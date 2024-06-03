package handlers

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/config"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/acme-sky/acmesky-api/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Struct used for login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Generate a valid JWT with HMAC 256 with an id and an expiration time of
// 1 hour. Key is stored in env.
func generateJWT(id uint) (string, error) {
	config, err := config.GetConfig()
	if err != nil {
		return "", err
	}

	key := []byte(config.String("jwt.token"))
	expiration := time.Now().Add(time.Hour)
	claims := &middleware.Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)

}

// Handler used to login the system and get a JWT to make requests.
// Password is stored as SHA256 hashed in database.
// Login godoc
//
//	@Summary	Make login
//	@Schemes
//	@Description	Make login
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/v1/login/ [post]
func LoginHandler(c *gin.Context) {
	db, err := db.GetDb()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(input.Password)))

	if err := db.Where("username = ? and password = ?", input.Username, password).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	token, err := generateJWT(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"user_id": user.Id,
	})
}
