package config

import (
	"github.com/fathoor/go-modular/internal/app/exception"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func NewFiber() *fiber.App {
	config := fiber.Config{
		ErrorHandler: exception.Handler, // Custom exception handler
		JSONEncoder:  json.Marshal,      // Use go-json for JSON encoding
		JSONDecoder:  json.Unmarshal,    // Use go-json for JSON decoding
	}
	app := fiber.New(config)
	app.Use(cors.New())    // Enable CORS
	app.Use(recover.New()) // Recover panics outside fiber

	log := logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} [${latency}]\n",
	}
	app.Use(logger.New(log)) // Log requests [Only for development, remove in production]

	// Default Route
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Endpoint `/` is not set.",
		})
	})

	// Health Check
	app.Get("/healthz", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "healthy",
		})
	})

	return app
}
