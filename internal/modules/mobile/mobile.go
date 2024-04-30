package mobile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvideMobile(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	homeRepository := postgres.NewHomeRepository(db)
	homeUseCase := usecase.NewHomeUseCase(&homeRepository)
	homeController := controller.NewHomeController(homeUseCase)

	profileRepository := postgres.NewProfileRepository(db)
	profileUseCase := usecase.NewProfileUseCase(&profileRepository)
	profileController := controller.NewProfileController(profileUseCase, validator)

	router.Route(
		app,
		homeController,
		profileController,
	)
}
