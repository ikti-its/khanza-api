package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/controller"
)

func Route(
	app *fiber.App,
	medisController *controller.MedisController,
) {
	inventaris := app.Group("/v1/inventaris")

	medis := inventaris.Group("/medis")
	{
		medis.Post("/", middleware.Authenticate([]int{1337, 1}), medisController.Create)
		medis.Get("/", middleware.Authenticate([]int{1337, 1, 2}), medisController.Get)
		medis.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), medisController.GetById)
		medis.Put("/:id", middleware.Authenticate([]int{1337, 1}), medisController.Update)
		medis.Delete("/:id", middleware.Authenticate([]int{1337, 1}), medisController.Delete)
	}
}
