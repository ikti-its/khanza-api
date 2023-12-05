package main

import (
	"github.com/fathoor/simkes-api/core/config"
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/provider"
	"github.com/gofiber/contrib/swagger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp(cfg)
		db  = config.ProvideDB(cfg)
		api = swagger.New(*config.ProvideSwagger())
	)

	provider.ProvideModule(app, db)
	app.Use(api)

	err := app.Listen(cfg.Get("APP_ADDRESS"))
	exception.PanicIfError(err)
}
