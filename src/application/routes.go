package application

import (
	"github.com/gofiber/fiber"

	"github.com/medson/superheroesAPI/src/handlers"
	"github.com/medson/superheroesAPI/src/helpers"
)

// SetupRoutes its responsible make the application routes available
//
// app - Represents the application
//
func SetupRoutes(app *fiber.App, handler handlers.Handler) {
	api := app.Group("/api/v1")

	api.Get("/healthCheck", handler.Status)

	api.Get("/supers", handler.Index)
	api.Post("/supers", handler.Create)
	api.Delete("/supers/:id", handler.Destroy)
	api.Get("/docs", helpers.ShowDocs)
}
