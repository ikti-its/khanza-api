package router

import (
	"github.com/gofiber/fiber/v2"
	dokterController "github.com/ikti-its/khanza-api/internal/modules/resep/internal/controller"
	obatController "github.com/ikti-its/khanza-api/internal/modules/resep/internal/controller"
)

func RegisterResepRoutes(app *fiber.App, resepObatController *obatController.ResepObatController, resepDokterController *dokterController.ResepDokterController) {
	// 🧾 Master route: Resep Obat
	resepObat := app.Group("/v1/resep-obat")
	resepObat.Get("/by-nomor-rawat/:nomor_rawat", resepObatController.GetByNomorRawat)
	resepObat.Post("/", resepObatController.Create)
	resepObat.Get("/", resepObatController.GetAll)
	resepObat.Get("/:no_resep", resepObatController.GetByNoResep)
	resepObat.Put("/", resepObatController.Update)
	resepObat.Delete("/:no_resep", resepObatController.Delete)
	resepObat.Put("/:no_resep/validasi", resepObatController.UpdateValidasi)

	// 💊 Detail route: Resep Dokter (resep_obat_detail)
	resepDokter := app.Group("/v1/resep-dokter")
	resepDokter.Post("/", resepDokterController.Create)
	resepDokter.Get("/", resepDokterController.GetAll)
	resepDokter.Get("/:no_resep", resepDokterController.GetByNoResep)
	resepDokter.Put("/", resepDokterController.Update)
	resepDokter.Delete("/:no_resep/:kode_barang", resepDokterController.Delete)
}
