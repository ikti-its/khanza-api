package router

import (
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller *controller.AuthController) {
	auth := app.Group("/v1/auth")
	{
		auth.Post("/login", controller.Login)
	}
}
