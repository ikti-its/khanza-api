package config

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber() *fiber.App {
	config := fiber.Config{
		ErrorHandler: exception.Handler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}
	app := fiber.New(config)
	app.Use(cors.New())    // Enable CORS
	app.Use(recover.New()) // Recover panics outside fiber

	log := logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} [${latency}]\n",
	}
	app.Use(logger.New(log)) // Log requests [Only for development, remove in production]

	return app
}
