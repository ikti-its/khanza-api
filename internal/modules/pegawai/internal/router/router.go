package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/pegawai/internal/controller"
)

func Route(
	app *fiber.App,
	pegawaiController *controller.PegawaiController,
	berkasController *controller.BerkasController,
	fotoController *controller.FotoController,
) {
	base := app.Group("/v1/pegawai")

	berkas := base.Group("/berkas")
	{
		berkas.Post("/", middleware.Authenticate([]int{1337, 1, 2}), berkasController.Create)
		berkas.Get("/", middleware.Authenticate([]int{1337, 1, 2}), berkasController.Get)
		berkas.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), berkasController.GetById)
		berkas.Put("/:id", middleware.Authenticate([]int{1337, 1, 2}), berkasController.Update)
		berkas.Delete("/:id", middleware.Authenticate([]int{1337, 1}), berkasController.Delete)
	}

	foto := base.Group("/foto")
	{
		foto.Post("/", middleware.Authenticate([]int{1337, 1, 2}), fotoController.Create)
		foto.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), fotoController.GetById)
		foto.Put("/:id", middleware.Authenticate([]int{1337, 1, 2}), fotoController.Update)
		foto.Delete("/:id", middleware.Authenticate([]int{1337, 1}), fotoController.Delete)
	}

	pegawai := base.Group("/")
	{
		pegawai.Post("/", middleware.Authenticate([]int{1337, 1}), pegawaiController.Create)
		pegawai.Get("/", middleware.Authenticate([]int{1337, 1, 2}), pegawaiController.Get)
		pegawai.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), pegawaiController.GetById)
		pegawai.Put("/:id", middleware.Authenticate([]int{1337, 1}), pegawaiController.Update)
		pegawai.Delete("/:id", middleware.Authenticate([]int{1337, 1}), pegawaiController.Delete)
	}
}
