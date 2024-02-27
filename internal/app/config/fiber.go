package config

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/gofiber/fiber/v2"
)

func ProvideFiber(cfg Config) *fiber.Config {
	return &fiber.Config{
		AppName:       cfg.Get("APP_NAME"),
		CaseSensitive: true,
		StrictRouting: false,
		ErrorHandler:  exception.Handler,
	}
}
