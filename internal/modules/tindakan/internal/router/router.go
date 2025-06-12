package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/controller"
)

func TindakanRoute(app *fiber.App, tindakanController *controller.TindakanController) {
	tindakan := app.Group("/v1/tindakan")

	tindakan.Get("/jenis", tindakanController.GetAllJenis)
	tindakan.Get("/jenis2", tindakanController.GetJenisByKodeQuery)
	tindakan.Get("/jenis/:kode", tindakanController.GetJenisByKode)
	tindakan.Post("/", tindakanController.Create)
	tindakan.Get("/", tindakanController.GetAll)
	tindakan.Get("/:nomor_rawat", tindakanController.GetByNomorRawat)
	tindakan.Put("/:nomor_rawat/:jam_rawat", tindakanController.Update)
	tindakan.Delete("/:nomor_rawat/:jam_rawat", tindakanController.Delete)

}
