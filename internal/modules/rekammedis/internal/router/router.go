package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/controller"
)

func PemeriksaanRanapRoute(app *fiber.App, pemeriksaanRanapController *controller.PemeriksaanRanapController) {
	// Group for the "pemeriksaanranap" routes under the "/v1/pemeriksaanranap" path
	pemeriksaanRanap := app.Group("/v1/pemeriksaanranap")

	pemeriksaanRanap.Get("/", pemeriksaanRanapController.GetAll)

	// Route for retrieving a specific pemeriksaan rawat inap by nomor_rawat
	pemeriksaanRanap.Get("/:nomor_rawat", pemeriksaanRanapController.GetByNomorRawat)

	// Route for creating a new pemeriksaan rawat inap
	pemeriksaanRanap.Post("/", pemeriksaanRanapController.Create)

	// Route for updating a specific pemeriksaan rawat inap by nomor_rawat
	pemeriksaanRanap.Put("/:nomor_rawat", pemeriksaanRanapController.Update)

	// Route for deleting a specific pemeriksaan rawat inap by nomor_rawat
	pemeriksaanRanap.Delete("/:nomor_rawat", pemeriksaanRanapController.Delete)

	// Additional routes specific to your requirements can be added here
}
