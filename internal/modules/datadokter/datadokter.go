package datadokter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/repository"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/datadokter/internal/usecase"
)

func ProvideDataDokter(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	repo := repository.NewRepository(db)
	uc := usecase.NewUseCase(repo)
	ctrl := controller.NewController(uc)

	router.Route(app, ctrl)
}
