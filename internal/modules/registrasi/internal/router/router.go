package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/controller"
)

func RegistrasiRoute(app *fiber.App, registrasiController *controller.RegistrasiController) {
	registrasi := app.Group("/v1/registrasi")

	registrasi.Get("/by-nomor-rm/:nomor_rm", registrasiController.GetByNomorRM)
	registrasi.Get("/by-no-rawat/:no_rawat", registrasiController.GetByNoRawat)
	registrasi.Put("/:nomor_reg/assign-room/:status", registrasiController.AssignRoomStatus)
	registrasi.Put("/status_kamar/:nomor_reg", registrasiController.UpdateStatusKamar)
	registrasi.Get("/pending-room", registrasiController.GetPendingRoomRequests)
	registrasi.Get("/dokter", registrasiController.GetAllDokter)
	registrasi.Post("/", registrasiController.Create)
	registrasi.Post("/submittambah", registrasiController.Create)
	registrasi.Get("/", registrasiController.GetAll)
	registrasi.Get("/:nomor_reg", registrasiController.GetByNomorReg)
	registrasi.Put("/:nomor_reg", registrasiController.Update)
	registrasi.Delete("/:nomor_reg", registrasiController.Delete)
	registrasi.Put("/:nomor_reg/assign-kamar", registrasiController.AssignKamar)

}
