package config

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ikti-its/khanza-api/internal/app/exception"
)

func NewFiber(cfg *Config) *fiber.App {
	config := fiber.Config{
		ErrorHandler: exception.Handler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}
	app := fiber.New(config)
	app.Use(cors.New())    // Enable CORS
	app.Use(recover.New()) // Recover panics outside fiber

	app.Get("/", func(ctx *fiber.Ctx) error { // Home Route
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Endpoint `/` is not set. Please refer to the API documentation at https://documenter.getpostman.com/view/23649536/2sA2rDy1iF",
		})
	})

	app.Use(healthcheck.New(healthcheck.Config{ // Health Check
		LivenessProbe: func(ctx *fiber.Ctx) bool {
			return true
		},
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

	if cfg.GetBool("APP_DDOS", false) { // Sinkhole if under-attack [Only for development, remove in production]
		app.Get("/play.mp3", func(ctx *fiber.Ctx) error { // Sinkhole Route
			return ctx.Redirect("https://www.youtube.com/watch?v=4hVhXKl-xyY", fiber.StatusMovedPermanently) // Rickroll :D
		})
	}

	return app
}
