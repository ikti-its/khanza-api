package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/controller"
)

func RegistrasiRoute(app *fiber.App, registrasiController *controller.RegistrasiController) {
	registrasi := app.Group("/v1/registrasi")

	registrasi.Post("/", registrasiController.Create)
	registrasi.Get("/", registrasiController.GetAll)
	registrasi.Get("/:nomor_reg", registrasiController.GetByNomorReg)
	registrasi.Put("/:nomor_reg", registrasiController.Update)
	registrasi.Delete("/:nomor_reg", registrasiController.Delete)
}
