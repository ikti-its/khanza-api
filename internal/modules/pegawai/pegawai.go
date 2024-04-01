package pegawai

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/controller"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/repository/postgres"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/router"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProvidePegawai(app *fiber.App, db *gorm.DB, validator *config.Validator) {
	pegawaiRepository := postgres.NewPegawaiRepository(db)
	pegawaiUseCase := usecase.NewPegawaiUseCase(&pegawaiRepository)
	pegawaiController := controller.NewPegawaiController(pegawaiUseCase, validator)

	berkasRepository := postgres.NewBerkasRepository(db)
	berkasUseCase := usecase.NewBerkasUseCase(&berkasRepository)
	berkasController := controller.NewBerkasController(berkasUseCase, validator)

	fotoRepository := postgres.NewFotoRepository(db)
	fotoUseCase := usecase.NewFotoUseCase(&fotoRepository)
	fotoController := controller.NewFotoController(fotoUseCase, validator)

	router.Route(
		app,
		pegawaiController,
		berkasController,
		fotoController,
	)
}
