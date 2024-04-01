package router

import (
	"github.com/fathoor/simkes-api/internal/app/middleware"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(
	app *fiber.App,
	pegawaiController *controller.PegawaiController,
	berkasController *controller.BerkasController,
	fotoController *controller.FotoController,
) {
	pegawai := app.Group("/v1/pegawai")
	{
		pegawai.Post("/", middleware.Authenticate([]int{1337, 1}), pegawaiController.Create)
		pegawai.Get("/", middleware.Authenticate([]int{1337, 1, 2}), pegawaiController.Get)
		pegawai.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), pegawaiController.GetById)
		pegawai.Put("/:id", middleware.Authenticate([]int{1337, 1}), pegawaiController.Update)
		pegawai.Delete("/:id", middleware.Authenticate([]int{1337, 1}), pegawaiController.Delete)
	}

	berkas := pegawai.Group("/berkas")
	{
		berkas.Post("/", middleware.Authenticate([]int{1337, 1, 2}), berkasController.Create)
		berkas.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), berkasController.GetById)
		berkas.Put("/:id", middleware.Authenticate([]int{1337, 1, 2}), berkasController.Update)
		berkas.Delete("/:id", middleware.Authenticate([]int{1337, 1}), berkasController.Delete)
	}

	foto := pegawai.Group("/foto")
	{
		foto.Post("/", middleware.Authenticate([]int{1337, 1, 2}), fotoController.Create)
		foto.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), fotoController.GetById)
		foto.Put("/:id", middleware.Authenticate([]int{1337, 1, 2}), fotoController.Update)
		foto.Delete("/:id", middleware.Authenticate([]int{1337, 1}), fotoController.Delete)
	}
}
