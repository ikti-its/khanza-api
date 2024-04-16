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

	router.Route(
		app,
		medisController,
	)
}
