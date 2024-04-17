package config

import (
	"github.com/fathoor/go-modular/internal/app/exception"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber(cfg *Config) *fiber.App {
	config := fiber.Config{
		ErrorHandler: exception.Handler, // Custom exception handler
		JSONEncoder:  json.Marshal,      // Use go-json for JSON encoding
		JSONDecoder:  json.Unmarshal,    // Use go-json for JSON decoding
	}
	app := fiber.New(config)
	app.Use(cors.New())    // Enable CORS
	app.Use(recover.New()) // Recover panics outside fiber

	app.Get("/", func(ctx *fiber.Ctx) error { // Home Route
		return ctx.Status(fiber.StatusNotFound).SendString("Endpoint / is not set.")
	})

	app.Use(healthcheck.New(healthcheck.Config{ // Health Check
		LivenessProbe:    func(ctx *fiber.Ctx) bool { return true },
		LivenessEndpoint: "/healthz",
	}))

	if cfg.GetBool("APP_DEBUG", false) { // Log requests [Only for development, remove in production]
		log := logger.Config{
			Format:     "[${time}] ${status} - ${method} ${path}\n", // e.g. [2006-01-02 15:04:05] 200 - GET /
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Jakarta",
		}
		app.Use(logger.New(log))
	}

	return app
}
