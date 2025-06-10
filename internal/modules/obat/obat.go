package obat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/repository"
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/usecase"
)

func ProvidePemberianObat(app *fiber.App, db *sqlx.DB) {
	var pemberianObatRepo repository.PemberianObatRepository = postgres.NewPemberianObatRepository(db)
	pemberianObatUseCase := usecase.NewPemberianObatUseCase(pemberianObatRepo)
	pemberianObatController := controller.NewPemberianObatController(pemberianObatUseCase)

	router.PemberianObatRoute(app, pemberianObatController)
}
