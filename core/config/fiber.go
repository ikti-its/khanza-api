package config

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/gofiber/fiber/v2"
)

func ProvideFiber(cfg Config) *fiber.Config {
	return &fiber.Config{
		AppName:       cfg.Get("APP_NAME"),
		Prefork:       cfg.GetBool("FIBER_PREFORK"),
		CaseSensitive: cfg.GetBool("FIBER_CASE_SENSITIVE"),
		StrictRouting: cfg.GetBool("FIBER_STRICT_ROUTING"),
		ErrorHandler:  exception.Handler,
	}
}
