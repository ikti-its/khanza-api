package rekammedis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvideRekamMedis(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	// Change repository to the PemeriksaanRanapRepository interface
	var pemeriksaanRanapRepository repository.PemeriksaanRanapRepository = postgres.NewPemeriksaanRanapRepository(db)

	// Use the PemeriksaanRanapUseCase instead of RegistrasiUseCase
	pemeriksaanRanapUseCase := usecase.NewPemeriksaanRanapUseCase(pemeriksaanRanapRepository)

	// Controller for PemeriksaanRanap
	pemeriksaanRanapController := controller.NewPemeriksaanRanapController(pemeriksaanRanapUseCase)

	// Ensure the function name matches the one in router.go
	router.PemeriksaanRanapRoute(app, pemeriksaanRanapController)
}
