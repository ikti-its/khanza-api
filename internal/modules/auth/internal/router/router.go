package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/controller"
)

func Route(app *fiber.App, controller *controller.AuthController) {
	auth := app.Group("/v1/auth")
	{
		auth.Post("/login", controller.Login)
	}
}
