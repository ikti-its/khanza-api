package router

import (
	"github.com/gofiber/fiber/v2"
	dokterController "github.com/ikti-its/khanza-api/internal/modules/resep/internal/controller"
	obatController "github.com/ikti-its/khanza-api/internal/modules/resep/internal/controller"
	racikanController "github.com/ikti-its/khanza-api/internal/modules/resep/internal/controller"
)

func RegisterResepRoutes(
	app *fiber.App,
	resepObatController *obatController.ResepObatController,
	resepDokterController *dokterController.ResepDokterController,
	resepDokterRacikanController *racikanController.ResepDokterRacikanController,
	resepDokterRacikanDetailController *racikanController.ResepDokterRacikanDetailController,
) {
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

	// 🧪 Racikan route: Resep Dokter Racikan
	resepRacikan := app.Group("/v1/resep-dokter-racikan")
	resepRacikan.Post("/", resepDokterRacikanController.Create)
	resepRacikan.Get("/", resepDokterRacikanController.GetAll)
	resepRacikan.Get("/:no_resep", resepDokterRacikanController.GetByNoResep)
	resepRacikan.Put("/", resepDokterRacikanController.Update)
	resepRacikan.Delete("/:no_resep/:no_racik", resepDokterRacikanController.Delete)

	// 🧬 Racikan Detail route: Resep Dokter Racikan Detail
	racikanDetail := app.Group("/v1/resep-dokter-racikan-detail")
	racikanDetail.Post("/", resepDokterRacikanDetailController.Create)
	racikanDetail.Get("/", resepDokterRacikanDetailController.GetAll)
	racikanDetail.Get("/:no_resep/:no_racik", resepDokterRacikanDetailController.GetByNoResepAndNoRacik)
	racikanDetail.Put("/", resepDokterRacikanDetailController.Update)
	racikanDetail.Delete("/:no_resep/:no_racik/:kode_brng", resepDokterRacikanDetailController.Delete)
}
