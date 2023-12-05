package config

import "github.com/gofiber/contrib/swagger"

func ProvideSwagger() *swagger.Config {
	return &swagger.Config{
		Next:     nil,
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/",
		Title:    "SIMKES RESTful API Documentation",
	}
}
