package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/trivia/pkg/handlers"
)

func NewRouter(logger handlers.Logger, db handlers.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		GETOnly: true,
	})
	app.Get("/api/v1/trivia", handlers.NewGetTriviaHandler(logger, db))
	return app
}
