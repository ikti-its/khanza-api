package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/controller"
)

func RekamMedisRoute(
	app *fiber.App,
	pemeriksaanRanapController *controller.PemeriksaanRanapController,
	catatanObservasiController *controller.CatatanObservasiRanapKebidananController,
	catatanPostpartumController *controller.CatatanObservasiRanapPostpartumController,
	catatanRanapController *controller.CatatanObservasiRanapController,
	diagnosaPasienController *controller.DiagnosaPasienController,
	resumePasienRanapController *controller.ResumePasienRanapController,
) {
	// ===== Pemeriksaan Ranap =====
	pemeriksaanRanap := app.Group("/v1/pemeriksaanranap")
	pemeriksaanRanap.Get("/", pemeriksaanRanapController.GetAll)
	pemeriksaanRanap.Get("/:nomor_rawat", pemeriksaanRanapController.GetByNomorRawat)
	pemeriksaanRanap.Post("/", pemeriksaanRanapController.Create)
	pemeriksaanRanap.Put("/:nomor_rawat", pemeriksaanRanapController.Update)
	pemeriksaanRanap.Delete("/:nomor_rawat", pemeriksaanRanapController.Delete)

	// ===== Catatan Observasi Ranap Kebidanan =====
	catatan := app.Group("/v1/catatan-observasi-ranap-kebidanan")
	catatan.Get("/", catatanObservasiController.GetAll)
	catatan.Get("/:no_rawat", catatanObservasiController.GetByNoRawat)
	catatan.Post("/", catatanObservasiController.Create)
	catatan.Put("/:no_rawat", catatanObservasiController.Update)
	catatan.Delete("/:no_rawat", catatanObservasiController.Delete)

	// ===== Catatan Observasi Ranap Postpartum =====
	postpartum := app.Group("/v1/catatan-observasi-ranap-postpartum")
	postpartum.Get("/", catatanPostpartumController.GetAll)
	postpartum.Get("/:no_rawat", catatanPostpartumController.GetByNoRawat)
	postpartum.Post("/", catatanPostpartumController.Create)
	postpartum.Put("/:no_rawat", catatanPostpartumController.Update)
	postpartum.Delete("/:no_rawat", catatanPostpartumController.Delete)

	// ===== Catatan Observasi Ranap (Umum) =====
	ranap := app.Group("/v1/catatan-observasi-ranap")
	ranap.Get("/", catatanRanapController.GetAll)
	ranap.Get("/:no_rawat", catatanRanapController.GetByNoRawat)
	ranap.Post("/", catatanRanapController.Create)
	ranap.Put("/:no_rawat", catatanRanapController.Update)
	ranap.Delete("/:no_rawat", catatanRanapController.Delete)

	// ===== Diagnosa Pasien =====
	diagnosa := app.Group("/v1/diagnosa-pasien")
	diagnosa.Get("/", diagnosaPasienController.GetAll)
	diagnosa.Get("/:no_rawat", diagnosaPasienController.GetByNoRawat)
	diagnosa.Get("/:no_rawat/:status", diagnosaPasienController.GetByNoRawatAndStatus)
	diagnosa.Post("/", diagnosaPasienController.Create)
	diagnosa.Put("/", diagnosaPasienController.Update)
	diagnosa.Delete("/:no_rawat/:kd_penyakit", diagnosaPasienController.Delete)

	// ===== Resume Pasien Ranap =====
	resume := app.Group("/v1/resume-pasien-ranap")
	resume.Get("/", resumePasienRanapController.GetAll)
	resume.Get("/:no_rawat", resumePasienRanapController.GetByNoRawat)
	resume.Post("/", resumePasienRanapController.Create)
	resume.Put("/:no_rawat", resumePasienRanapController.Update)
	resume.Delete("/:no_rawat", resumePasienRanapController.Delete)
}
