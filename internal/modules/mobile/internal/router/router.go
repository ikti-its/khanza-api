package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/controller"
)

func Route(
	app *fiber.App,
	homeController *controller.HomeController,
	profileController *controller.ProfileController,
	pegawaiController *controller.PegawaiController,
	ketersediaanController *controller.KetersediaanController,
	kehadiranController *controller.KehadiranController,
	jadwalController *controller.JadwalController,
) {
	mobile := app.Group("/v1/m")

	home := mobile.Group("/home")
	{
		home.Get("/", middleware.Authenticate([]int{0}), homeController.GetHome)
	}

	profile := mobile.Group("/profile")
	{
		profile.Put("/:id", middleware.Authenticate([]int{0}), profileController.Update)
	}

	pegawai := mobile.Group("/pegawai")
	{
		pegawai.Get("/", middleware.Authenticate([]int{0}), pegawaiController.GetById)
	}

	ketersediaan := mobile.Group("/ketersediaan")
	{
		ketersediaan.Get("/", middleware.Authenticate([]int{1337, 1}), ketersediaanController.Get)
	}

	kehadiran := mobile.Group("/kehadiran")
	{
		kehadiran.Get("/jadwal/:id", middleware.Authenticate([]int{0}), jadwalController.GetByPegawaiId)
		kehadiran.Get("/:id", middleware.Authenticate([]int{0}), kehadiranController.GetByPegawaiId)
	}
}
