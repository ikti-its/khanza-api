package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/dokterjaga/internal/controller"
)

func DokterJagaRoute(app *fiber.App, dokterJagaController *controller.DokterJagaController) {
	dokter := app.Group("/v1/dokterjaga")

	dokter.Post("/", dokterJagaController.Create)
	dokter.Get("/", dokterJagaController.GetAll)
	dokter.Get("/:kode_dokter", dokterJagaController.GetByKodeDokter)
	dokter.Put("/", dokterJagaController.Update)
	dokter.Delete("/:kode_dokter", dokterJagaController.Delete) // expects query param ?hari_kerja=YYYY-MM-DD

	dokter.Get("/status/:status", dokterJagaController.GetByStatus)
	dokter.Put("/update-status", dokterJagaController.UpdateStatus)
}
