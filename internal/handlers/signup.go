package handlers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Struct used for signup
type SignupInput struct {
	Username           string `binding:"required"`
	Password           string `binding:"required"`
	Name               string `binding:"required"`
	Email              string `binding:"required"`
	Address            *string
	ProntogramUsername *string
}

func (in SignupInput) GetUsername() string {
	return in.Username
}

func (in SignupInput) GetPassword() string {
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(in.Password)))
	return password
}

func (in SignupInput) GetName() string {
	return in.Name
}

func (in SignupInput) GetEmail() string {
	return in.Email
}

func (in SignupInput) GetAddress() *string {
	return in.Address
}

func (in SignupInput) GetProntogramUsername() *string {
	return in.ProntogramUsername
}

func (in SignupInput) Validate(db *gorm.DB) error {
	var user *models.User
	if err := db.Where("username = ? OR email = ?", in.GetUsername(), in.GetEmail()).First(&user).Error; err == nil {
		return errors.New("user with this username or email already exists")
	}

	const emailRegexPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)

	if !re.MatchString(in.GetEmail()) {
		return errors.New("this email has a bad format")
	}

	if len(in.GetPassword()) < 8 {
		return errors.New("choose a password with a length of at least 8 chars")
	}

	return nil
}

// Handler used to signup to the system. In other words: create a new user.
// Signup godoc
//
//	@Summary	Make signup
//	@Schemes
//	@Description	Make signup
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		201
//	@Router			/v1/signup/ [post]
func SignupHandler(c *gin.Context) {
	db, err := db.GetDb()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := input.Validate(db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := models.NewUser(input)
	db.Create(&user)
	db.Omit("Password").First(&user, user.Id)

	c.JSON(http.StatusCreated, user)
}
