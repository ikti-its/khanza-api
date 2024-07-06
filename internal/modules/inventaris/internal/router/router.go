package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/middleware"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/controller"
)

func Route(
	app *fiber.App,
	medisController *controller.MedisController,
	obatController *controller.ObatController,
	alkesController *controller.AlkesController,
	bhpController *controller.BhpController,
	darahController *controller.DarahController,
	stokController *controller.StokController,
	supplierController *controller.SupplierController,
	satuanController *controller.SatuanController,
	transaksiController *controller.TransaksiController,
) {
	inventaris := app.Group("/v1/inventaris")

	medis := inventaris.Group("/medis")
	{
		medis.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), medisController.Create)
		medis.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), medisController.Get)
		medis.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), medisController.GetById)
		medis.Put("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), medisController.Update)
		medis.Delete("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), medisController.Delete)
	}

	obat := inventaris.Group("/obat")
	{
		obat.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), obatController.Create)
		obat.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), obatController.Get)
		obat.Get("/medis/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), obatController.GetByIdMedis)
		obat.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), obatController.GetById)
		obat.Put("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), obatController.Update)
		obat.Delete("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), obatController.Delete)
	}

	alkes := inventaris.Group("/alkes")
	{
		alkes.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), alkesController.Create)
		alkes.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), alkesController.Get)
		alkes.Get("/medis/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), alkesController.GetByIdMedis)
		alkes.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), alkesController.GetById)
		alkes.Put("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), alkesController.Update)
		alkes.Delete("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), alkesController.Delete)
	}

	bhp := inventaris.Group("/bhp")
	{
		bhp.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), bhpController.Create)
		bhp.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), bhpController.Get)
		bhp.Get("/medis/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), bhpController.GetByIdMedis)
		bhp.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), bhpController.GetById)
		bhp.Put("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), bhpController.Update)
		bhp.Delete("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), bhpController.Delete)
	}

	darah := inventaris.Group("/darah")
	{
		darah.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), darahController.Create)
		darah.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), darahController.Get)
		darah.Get("/medis/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), darahController.GetByIdMedis)
		darah.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), darahController.GetById)
		darah.Put("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), darahController.Update)
		darah.Delete("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002}), darahController.Delete)
	}

	stok := inventaris.Group("/stok")
	{
		stok.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4004}), stokController.Create)
		stok.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), stokController.Get)
		stok.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), stokController.GetById)
		stok.Put("/:id", middleware.Authenticate([]int{1337, 4000, 4004}), stokController.Update)
		stok.Delete("/:id", middleware.Authenticate([]int{1337, 4000, 4004}), stokController.Delete)
	}

	supplier := inventaris.Group("/supplier")
	{
		supplier.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), supplierController.Get)
	}

	satuan := inventaris.Group("/satuan")
	{
		satuan.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), satuanController.Get)
	}

	transaksi := inventaris.Group("/transaksi")
	{
		transaksi.Post("/", middleware.Authenticate([]int{1337, 1, 4000, 4004}), transaksiController.Create)
		transaksi.Get("/", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), transaksiController.Get)
		transaksi.Get("/stok/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), transaksiController.GetByStokId)
		transaksi.Get("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4001, 4002, 4003, 4004}), transaksiController.GetById)
		transaksi.Put("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4004}), transaksiController.Update)
		transaksi.Delete("/:id", middleware.Authenticate([]int{1337, 1, 4000, 4004}), transaksiController.Delete)
	}
}
