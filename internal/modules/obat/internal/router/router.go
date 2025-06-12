package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/controller"
)

func PemberianObatRoute(app *fiber.App, pemberianObatController *controller.PemberianObatController) {
	obat := app.Group("/v1/pemberian-obat")

	obat.Get("/databarang", pemberianObatController.GetAllDataBarang)
	obat.Post("/", pemberianObatController.Create)
	obat.Get("/", pemberianObatController.GetAll)
	obat.Get("/:nomor_rawat", pemberianObatController.GetByNomorRawat)
	obat.Put("/:nomor_rawat", pemberianObatController.Update)
	obat.Delete("/:nomor_rawat/:jam_beri", pemberianObatController.Delete)
}

func GudangBarangRoute(app *fiber.App, gudangBarangController *controller.GudangBarangController) {
	gudang := app.Group("/v1/gudang-barang")

	gudang.Post("/", gudangBarangController.Create)
	gudang.Get("/", gudangBarangController.GetAll)
	gudang.Get("/:id", gudangBarangController.GetByID)
	gudang.Put("/:id", gudangBarangController.Update)
	gudang.Delete("/:id", gudangBarangController.Delete)
}
