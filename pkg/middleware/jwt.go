package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/acme-sky/acmesky-api/internal/models"
	"github.com/acme-sky/acmesky-api/pkg/config"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
)

// Claims for JWT. We store all the JWT default claims + username for this
// software.
type Claims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

// Check if a `bearer` token is valid or not
func checkAuth(bearer string) (*jwt.Token, error) {
	conf, err := config.GetConfig()

	if err != nil {
		return nil, err
	}

	key := []byte(conf.String("jwt.token"))

	// If header does not start with "Bearer " better to stop here
	if !strings.HasPrefix(bearer, "Bearer ") {
		return nil, errors.New("unauthorized")
	}

	// JWT is parsed only by the last part of the Authorization header
	token, err := jwt.ParseWithClaims(strings.Split(bearer, " ")[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("unauthorized")
		} else {
			return nil, errors.New("bad request")
		}
	} else if token == nil || !token.Valid {
		return nil, errors.New("unauthorized")
	}

	return token, nil
}

// Check the authorization from the header bearer token. If the authorization is
// good does nothing, else it aborts the Gin context.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := checkAuth(c.Request.Header.Get("Authorization"))
		if err != nil {
			switch err.Error() {
			case "unauthorized":
				c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				break
			case "bad request":
				c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
				break
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}

			c.Abort()
			return
		}

		c.Set("token", token)

		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("user_id", claims.Id)
		}
	}
}

// Check if the authorized user is an admin
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := checkAuth(c.Request.Header.Get("Authorization"))
		if err != nil {
			switch err.Error() {
			case "unauthorized":
				c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				break
			case "bad request":
				c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
				break
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}

			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok {
			db, _ := db.GetDb()
			if err := db.Where("id = ? and is_admin = true", claims.Id).First(&models.User{}).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				c.Abort()
				return
			}
			c.Set("user_id", claims.Id)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "can't parse this user"})
			c.Abort()
			return
		}

		c.Set("token", token)
	}
}
