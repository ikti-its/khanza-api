package test

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/provider"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func ProvideTestApp() *fiber.App {
	err := godotenv.Load(".env.test")
	exception.PanicIfError(err)

	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp(cfg)
		db  = config.ProvideDB(cfg)
		di  = provider.Provider{App: app, DB: db}
	)

	di.Provide()

	return app
}

var (
	app   = ProvideTestApp()
	token = os.Getenv("TEST_TOKEN")
)
