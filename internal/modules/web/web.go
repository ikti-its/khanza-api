package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/web/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/web/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/web/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/web/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvideWeb(app *fiber.App, db *sqlx.DB) {
	homeRepository := postgres.NewHomeRepository(db)
	homeUseCase := usecase.NewHomeUseCase(&homeRepository)
	homeController := controller.NewHomeController(homeUseCase)

	router.Route(
		app,
		homeController,
	)
}
