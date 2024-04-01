package main

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/provider"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	var (
		cfg       = config.NewConfig()
		fiber     = config.NewFiber()
		postgres  = config.NewPostgres(cfg)
		validator = config.NewValidator()
		bootstrap = provider.Provider{App: fiber, Config: cfg, DB: postgres, Validator: validator}
	)

	bootstrap.Provide()

	if err := fiber.Listen(fmt.Sprintf(":%d", cfg.GetInt("APP_PORT", 8080))); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
