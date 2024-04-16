package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/controller"
)

func Route(
	app *fiber.App,
	pengajuanController *controller.PengajuanController,
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
}
