package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/web/internal/controller"
)

func Route(
	app *fiber.App,
	homeController *controller.HomeController,
) {
	web := app.Group("/v1/w")

	home := web.Group("/home")
	{
		home.Get("/pegawai", middleware.Authenticate([]int{1337, 1, 2}), homeController.GetHomePegawai)
	}
}
