package config

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func ProvideApp(cfg Config) *fiber.App {
	var (
		app = fiber.New(*ProvideFiber(cfg))
		api = swagger.New(*ProvideSwagger()) // COMMENT WHILE TEST
		_   = ProvideDB(cfg)
	)

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(api) // COMMENT WHILE TEST

	return app
}
