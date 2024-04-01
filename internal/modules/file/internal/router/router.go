package router

import (
	"github.com/fathoor/simkes-api/internal/app/middleware"
	"github.com/fathoor/simkes-api/internal/modules/file/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller *controller.FileController) {
	file := app.Group("/v1/file")
	{
		file.Post("/:type", middleware.Authenticate([]int{0}), controller.Upload)
		file.Get("/:type/:name", middleware.Authenticate([]int{0}), controller.View)
		file.Get("/:type/:name/download", middleware.Authenticate([]int{0}), controller.Download)
		file.Delete("/:type/:name", middleware.Authenticate([]int{1337, 1}), controller.Delete)
	}
}
