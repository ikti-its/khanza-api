package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/controller"
)

func Route(
	app *fiber.App,
	medisController *controller.MedisController,
	obatController *controller.ObatController,
	alkesController *controller.AlkesController,
	bhpController *controller.BhpController,
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

	obat := inventaris.Group("/obat")
	{
		obat.Post("/", middleware.Authenticate([]int{1337, 1}), obatController.Create)
		obat.Get("/", middleware.Authenticate([]int{1337, 1, 2}), obatController.Get)
		obat.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), obatController.GetById)
		obat.Put("/:id", middleware.Authenticate([]int{1337, 1}), obatController.Update)
		obat.Delete("/:id", middleware.Authenticate([]int{1337, 1}), obatController.Delete)
	}

	alkes := inventaris.Group("/alkes")
	{
		alkes.Post("/", middleware.Authenticate([]int{1337, 1}), alkesController.Create)
		alkes.Get("/", middleware.Authenticate([]int{1337, 1, 2}), alkesController.Get)
		alkes.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), alkesController.GetById)
		alkes.Put("/:id", middleware.Authenticate([]int{1337, 1}), alkesController.Update)
		alkes.Delete("/:id", middleware.Authenticate([]int{1337, 1}), alkesController.Delete)
	}

	bhp := inventaris.Group("/bhp")
	{
		bhp.Post("/", middleware.Authenticate([]int{1337, 1}), bhpController.Create)
		bhp.Get("/", middleware.Authenticate([]int{1337, 1, 2}), bhpController.Get)
		bhp.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), bhpController.GetById)
		bhp.Put("/:id", middleware.Authenticate([]int{1337, 1}), bhpController.Update)
		bhp.Delete("/:id", middleware.Authenticate([]int{1337, 1}), bhpController.Delete)
	}
}
