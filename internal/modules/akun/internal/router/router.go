package router

import (
	"github.com/fathoor/simkes-api/internal/app/middleware"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(
	app *fiber.App,
	controller *controller.AkunController,
	alamatController *controller.AlamatController,
) {
	akun := app.Group("/v1/akun")
	{
		akun.Post("/", middleware.Authenticate([]int{1337, 1}), controller.Create)
		akun.Get("/", middleware.Authenticate([]int{1337, 1, 2}), controller.Get)
		akun.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), controller.GetById)
		akun.Put("/:id", middleware.Authenticate([]int{1337, 1, 2}), controller.Update)
		akun.Delete("/:id", middleware.Authenticate([]int{1337, 1}), controller.Delete)
	}

	alamat := akun.Group("/alamat")
	{
		alamat.Post("/", middleware.Authenticate([]int{0}), alamatController.Create)
		alamat.Get("/:id", middleware.Authenticate([]int{0}), alamatController.GetById)
		alamat.Put("/:id", middleware.Authenticate([]int{0}), alamatController.Update)
		alamat.Delete("/:id", middleware.Authenticate([]int{1337, 1}), alamatController.Delete)
	}
}
