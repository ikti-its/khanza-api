package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/controller"
)

func Route(
	app *fiber.App,
	pengajuanController *controller.PengajuanController,
	persetujuanController *controller.PersetujuanController,
	pesananController *controller.PesananController,
) {
	pengadaan := app.Group("/v1/pengadaan")

	pengajuan := pengadaan.Group("/pengajuan")
	{
		pengajuan.Post("/", middleware.Authenticate([]int{1337, 1}), pengajuanController.Create)
		pengajuan.Get("/", middleware.Authenticate([]int{1337, 1, 2}), pengajuanController.Get)
		pengajuan.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), pengajuanController.GetById)
		pengajuan.Put("/:id", middleware.Authenticate([]int{1337, 1}), pengajuanController.Update)
		pengajuan.Delete("/:id", middleware.Authenticate([]int{1337, 1}), pengajuanController.Delete)
	}

	persetujuan := pengadaan.Group("/persetujuan")
	{
		persetujuan.Post("/", middleware.Authenticate([]int{1337, 1}), persetujuanController.Create)
		persetujuan.Get("/", middleware.Authenticate([]int{1337, 1, 2}), persetujuanController.Get)
		persetujuan.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), persetujuanController.GetById)
		persetujuan.Put("/:id", middleware.Authenticate([]int{1337, 1}), persetujuanController.Update)
		persetujuan.Delete("/:id", middleware.Authenticate([]int{1337, 1}), persetujuanController.Delete)
	}

	pesanan := pengadaan.Group("/pesanan")
	{
		pesanan.Post("/", middleware.Authenticate([]int{1337, 1}), pesananController.Create)
		pesanan.Get("/", middleware.Authenticate([]int{1337, 1, 2}), pesananController.Get)
		pesanan.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), pesananController.GetById)
		pesanan.Put("/:id", middleware.Authenticate([]int{1337, 1}), pesananController.Update)
		pesanan.Delete("/:id", middleware.Authenticate([]int{1337, 1}), pesananController.Delete)
	}
}
