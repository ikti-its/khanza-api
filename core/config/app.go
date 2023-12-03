package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func ProvideApp(cfg Config) *fiber.App {
	var (
		app = fiber.New(*ProvideFiber(cfg))
		_   = ProvideDB(cfg)
	)

	app.Use(recover.New())
	app.Use(cors.New())

	return app
}
