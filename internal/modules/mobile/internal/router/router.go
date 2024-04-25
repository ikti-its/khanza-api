package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/controller"
)

func Route(
	app *fiber.App,
	homeController *controller.HomeController,
) {
	mobile := app.Group("/v1/m")

	home := mobile.Group("/home")
	{
		home.Get("/pegawai/:id", middleware.Authenticate([]int{1337, 1, 2}), homeController.GetHomePegawai)
	}
}
