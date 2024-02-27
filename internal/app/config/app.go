package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func ProvideApp(cfg Config) *fiber.App {
	app := fiber.New(*ProvideFiber(cfg))

	app.Use(recover.New())
	app.Use(cors.New())

	return app
}
