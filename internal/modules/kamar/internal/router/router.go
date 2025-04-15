package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/controller"
)

func KamarRoute(app *fiber.App, kamarController *controller.KamarController) {
	kamar := app.Group("/v1/kamar")

	kamar.Get("/available", kamarController.GetAvailable)
	kamar.Post("/", kamarController.Create)
	kamar.Get("/", kamarController.GetAll)
	kamar.Get("/:nomor_bed", kamarController.GetByNomorBed)
	kamar.Put("/:nomor_bed", kamarController.Update)
	kamar.Delete("/:nomor_bed", kamarController.Delete)

}
