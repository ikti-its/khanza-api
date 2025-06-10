package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/controller"
)

func RujukanRoute(
	app *fiber.App,
	masukCtrl *controller.RujukanMasukController,
	keluarCtrl *controller.RujukanKeluarController,
) {
	// Group for Rujukan Masuk
	rujukanMasuk := app.Group("/v1/rujukanmasuk")
	rujukanMasuk.Post("/", masukCtrl.Create)
	rujukanMasuk.Get("/", masukCtrl.GetAll)
	rujukanMasuk.Get("/:nomor_rawat", masukCtrl.GetByNomorRawat)
	rujukanMasuk.Put("/:nomor_rawat", masukCtrl.Update)
	rujukanMasuk.Delete("/:nomor_rawat", masukCtrl.Delete)

	// Group for Rujukan Keluar
	rujukanKeluar := app.Group("/v1/rujukankeluar")
	rujukanKeluar.Post("/", keluarCtrl.Create)
	rujukanKeluar.Get("/", keluarCtrl.GetAll)
	rujukanKeluar.Get("/:nomor_rawat", keluarCtrl.GetByNomorRawat)
	rujukanKeluar.Put("/:nomor_rawat", keluarCtrl.Update)
	rujukanKeluar.Delete("/:nomor_rawat", keluarCtrl.Delete)
}
