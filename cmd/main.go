package main

import (
	"log"

	"github.com/acme-sky/acmesky-api/internal/handlers"
	"github.com/acme-sky/acmesky-api/pkg/config"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/acme-sky/acmesky-api/pkg/middleware"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// Create a new instance of Gin server
func main() {
	router := gin.Default()

	var err error

	// Read environment variables and stops execution if any errors occur
	err = config.LoadConfig()
	if err != nil {
		log.Printf("failed to load config. err %v", err)

		return
	}

	// Ignore error because if it failed on loading, it should raised an error
	// above.
	config, _ := config.GetConfig()

	if _, err := db.InitDb(config.String("database.dsn")); err != nil {
		log.Printf("failed to connect database. err %v", err)

		return
	}

	// Env variable `debug` set up the mode below
	if !config.Bool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.Default())

	router.StaticFile("/swagger.yml", "cmd/swagger.yml")

	// v1 is just like a namespace for every routing here below
	v1 := router.Group("/v1")
	{
		v1.POST("/login/", handlers.LoginHandler)
		v1.POST("/signup/", handlers.SignupHandler)

		users := v1.Group("/users")
		{
			users.GET("/", middleware.Admin(), handlers.UserHandlerGet)
			users.GET("/:id/", middleware.OwnerOrAdmin(), handlers.UserHandlerGetId)
		}
	}

	router.Run()
}
