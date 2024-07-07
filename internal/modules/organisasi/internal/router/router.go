package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/controller"
)

func Route(
	app *fiber.App,
	organisasiController *controller.OrganisasiController,
) {
	organisasi := app.Group("/v1/organisasi")
	{
		organisasi.Post("/", middleware.Authenticate([]int{1337, 1}), organisasiController.Create)
		organisasi.Get("/", middleware.Authenticate([]int{0}), organisasiController.Get)
		organisasi.Get("/current", middleware.Authenticate([]int{0}), organisasiController.GetCurrent)
		organisasi.Get("/:id", middleware.Authenticate([]int{0}), organisasiController.GetById)
		organisasi.Put("/:id", middleware.Authenticate([]int{1337, 1}), organisasiController.Update)
		organisasi.Delete("/:id", middleware.Authenticate([]int{1337, 1}), organisasiController.Delete)
	}
}
