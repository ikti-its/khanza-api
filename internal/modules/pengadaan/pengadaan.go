package pengadaan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvidePengadaan(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	pengajuanRepository := postgres.NewPengajuanRepository(db)
	pengajuanUseCase := usecase.NewPengajuanUseCase(&pengajuanRepository)
	pengajuanController := controller.NewPengajuanController(pengajuanUseCase, validator)

	persetujuanRepository := postgres.NewPersetujuanRepository(db)
	persetujuanUseCase := usecase.NewPersetujuanUseCase(&persetujuanRepository)
	persetujuanController := controller.NewPersetujuanController(persetujuanUseCase, validator)

	router.Route(
		app,
		pengajuanController,
		persetujuanController,
	)
}
