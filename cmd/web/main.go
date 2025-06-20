package main

import (
	"log"

	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/app/provider"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg       = config.NewConfig()
		fiber     = config.NewFiber(cfg)
		postgres  = config.NewPostgres(cfg)
		validator = config.NewValidator()
		bootstrap = provider.Provider{App: fiber, Config: cfg, PG: postgres, Validator: validator}
	)

	defer func(postgres *sqlx.DB) {
		err := postgres.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(postgres)

	bootstrap.Provide()

	// routes := fiber.GetRoutes()
	// for _, route := range routes {
	// 	log.Printf("METHOD: %s | PATH: %s", route.Method, route.Path)
	// }

	// if err := fiber.Listen(fmt.Sprintf(":%d", cfg.GetInt("APP_PORT", 8080))); err != nil {
	if err := fiber.Listen("[::]:8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
