package mobile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvideMobile(app *fiber.App, db *sqlx.DB) {
	homeRepository := postgres.NewHomeRepository(db)
	homeUseCase := usecase.NewHomeUseCase(&homeRepository)
	homeController := controller.NewHomeController(homeUseCase)

	router.Route(
		app,
		homeController,
	)
}
