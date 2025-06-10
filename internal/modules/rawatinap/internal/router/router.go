package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/rawatinap/internal/controller"
)

func RawatInapRoute(app *fiber.App, rawatInapController *controller.RawatInapController) {
	rawat := app.Group("/v1/rawatinap")

	rawat.Post("/", rawatInapController.Create)
	rawat.Get("/", rawatInapController.GetAll)
	rawat.Get("/:nomor_rawat", rawatInapController.GetByNomorRawat)
	rawat.Put("/:nomor_rawat", rawatInapController.Update)
	rawat.Delete("/:nomor_rawat", rawatInapController.Delete)
}
