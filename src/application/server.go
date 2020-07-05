package application

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/medson/superheroesAPI/src/handlers"
	"github.com/medson/superheroesAPI/src/infra"
)

// NewServer its responsible start a server instance
//
// Returns a pointer with a fiber app
//
func NewServer() *fiber.App {
	app := fiber.New()
	if err := infra.InitDB(); err != nil {
		log.Fatalf("Error in intializing database: %s", err)
	}

	SetupRoutes(app, handlers.Handler{})
	app.Use("/swagger", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:3000/api/v1/docs",
		DeepLinking: false,
	}))
	return app
}
