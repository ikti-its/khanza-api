package main

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/provider"
	"github.com/gofiber/contrib/swagger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp(cfg)
		db  = config.ProvideDB(cfg)
		di  = provider.Provider{App: app, DB: db}
		api = swagger.New(*config.ProvideSwagger())
	)

	di.Provide()
	app.Use(api)

	err := app.Listen(fmt.Sprintf(":%d", cfg.GetInt("APP_PORT")))
	exception.PanicIfError(err)
}
