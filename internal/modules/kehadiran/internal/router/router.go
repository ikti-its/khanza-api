package router

import (
	"github.com/fathoor/simkes-api/internal/app/middleware"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(
	app *fiber.App,
	kehadiranController *controller.KehadiranController,
	jadwalController *controller.JadwalController,
) {
	kehadiran := app.Group("/v1/kehadiran")
	{
		kehadiran.Post("/attend", middleware.Authenticate([]int{1337, 1, 2}), kehadiranController.Attend)
		kehadiran.Post("/leave", middleware.Authenticate([]int{1337, 1, 2}), kehadiranController.Leave)
		kehadiran.Get("/", middleware.Authenticate([]int{1337, 1}), kehadiranController.Get)
		kehadiran.Get("/pegawai/:id", middleware.Authenticate([]int{1337, 1, 2}), kehadiranController.GetByPegawaiId)
		kehadiran.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), kehadiranController.GetById)
		kehadiran.Put("/:id", middleware.Authenticate([]int{1337, 1}), kehadiranController.Update)
	}

	jadwal := kehadiran.Group("/jadwal")
	{
		jadwal.Get("/", middleware.Authenticate([]int{1337, 1}), jadwalController.Get)
		jadwal.Get("/hari/:id", middleware.Authenticate([]int{1337, 1}), jadwalController.GetByHariId)
		jadwal.Get("/pegawai/:id", middleware.Authenticate([]int{1337, 1, 2}), jadwalController.GetByPegawaiId)
		jadwal.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), jadwalController.GetById)
		jadwal.Put("/:id", middleware.Authenticate([]int{1337, 1}), jadwalController.Update)
	}
}
