package main

import (
	"log"

	"github.com/acme-sky/acmesky-api/internal/handlers"
	"github.com/acme-sky/acmesky-api/pkg/config"
	"github.com/acme-sky/acmesky-api/pkg/db"
	"github.com/acme-sky/acmesky-api/pkg/message"
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

	if err := message.CreateConnection(config.String("rabbitmq")); err != nil {
		log.Printf("failed to create a connection to RabbitMQ. err %v", err)

		return
	}

	// Env variable `debug` set up the mode below
	if !config.Bool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.AllowAll())

	router.StaticFile("/swagger.yml", "cmd/swagger.yml")

	// v1 is just like a namespace for every routing here below
	v1 := router.Group("/v1")
	{
		v1.POST("/login/", handlers.LoginHandler)
		v1.POST("/signup/", handlers.SignupHandler)

		users := v1.Group("/users")
		{
			users.GET("/", middleware.Admin(), handlers.UserHandlerGet)
			users.GET("/:id/", middleware.Auth(), handlers.UserHandlerGetId)
			users.PUT("/:id/", middleware.Auth(), handlers.UserHandlerPut)
		}

		interests := v1.Group("/interests")
		{
			interests.Use(middleware.Auth())
			interests.GET("/", handlers.InterestHandlerGet)
			interests.POST("/", handlers.InterestHandlerPost)
			interests.GET("/:id/", handlers.InterestHandlerGetId)
			interests.DELETE("/:id/", handlers.InterestHandlerDelete)
		}

		availableFlights := v1.Group("/available-flights")
		{
			availableFlights.Use(middleware.Auth())
			availableFlights.GET("/", handlers.AvailableFlightHandlerGet)
			availableFlights.GET("/:id/", handlers.AvailableFlightHandlerGetId)
		}

		journeys := v1.Group("/journeys")
		{
			journeys.Use(middleware.Auth())
			journeys.GET("/", handlers.JourneyHandlerGet)
			journeys.GET("/:id/", handlers.JourneyHandlerGetId)
		}

		offers := v1.Group("/offers")
		{
			offers.GET("/", middleware.Auth(), handlers.OfferHandlerGet)
			offers.GET("/:id/", middleware.Auth(), handlers.OfferHandlerGetId)
			offers.POST("/confirm/", middleware.Auth(), handlers.OfferConfirmHandlerPost)
			offers.POST("/pay/:id/", handlers.OfferHandlerPay)
		}

		invoices := v1.Group("/invoices")
		{
			invoices.Use(middleware.Auth())
			invoices.GET("/", handlers.InvoiceHandlerGet)
			invoices.GET("/:id/", handlers.InvoiceHandlerGetId)
		}
	}

	router.Run(config.String("server.url"))
	message.CloseConnection()
}
