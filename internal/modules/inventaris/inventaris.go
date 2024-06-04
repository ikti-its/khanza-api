package inventaris

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvideInventaris(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	medisRepository := postgres.NewMedisRepository(db)
	medisUseCase := usecase.NewMedisUseCase(&medisRepository)
	medisController := controller.NewMedisController(medisUseCase, validator)

	obatRepository := postgres.NewObatRepository(db)
	obatUseCase := usecase.NewObatUseCase(&obatRepository)
	obatController := controller.NewObatController(obatUseCase, validator)

	alkesRepository := postgres.NewAlkesRepository(db)
	alkesUseCase := usecase.NewAlkesUseCase(&alkesRepository)
	alkesController := controller.NewAlkesController(alkesUseCase, validator)

	bhpRepository := postgres.NewBhpRepository(db)
	bhpUseCase := usecase.NewBhpUseCase(&bhpRepository)
	bhpController := controller.NewBhpController(bhpUseCase, validator)

	darahRepository := postgres.NewDarahRepository(db)
	darahUseCase := usecase.NewDarahUseCase(&darahRepository)
	darahController := controller.NewDarahController(darahUseCase, validator)

	stokRepository := postgres.NewStokRepository(db)
	stokUseCase := usecase.NewStokUseCase(&stokRepository)
	stokController := controller.NewStokController(stokUseCase, validator)

	supplierRepository := postgres.NewSupplierRepository(db)
	supplierUseCase := usecase.NewSupplierUseCase(&supplierRepository)
	supplierController := controller.NewSupplierController(supplierUseCase)

	satuanRepository := postgres.NewSatuanRepository(db)
	satuanUseCase := usecase.NewSatuanUseCase(&satuanRepository)
	satuanController := controller.NewSatuanController(satuanUseCase)

	transaksiRepository := postgres.NewTransaksiRepository(db)
	transaksiUseCase := usecase.NewTransaksiUseCase(&transaksiRepository)
	transaksiController := controller.NewTransaksiController(transaksiUseCase, validator)

	router.Route(
		app,
		medisController,
		obatController,
		alkesController,
		bhpController,
		darahController,
		stokController,
		supplierController,
		satuanController,
		transaksiController,
	)
}
