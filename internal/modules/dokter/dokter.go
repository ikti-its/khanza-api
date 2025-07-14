package dokter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/dokter/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/dokter/internal/repository"
	"github.com/ikti-its/khanza-api/internal/modules/dokter/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/dokter/internal/usecase"
)

func ProvideDokter(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	repo := repository.NewRepository(db)
	uc := usecase.NewUseCase(repo)
	ctrl := controller.NewController(uc)

	router.Route(app, ctrl)
}
