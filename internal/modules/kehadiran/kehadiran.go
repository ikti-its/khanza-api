package kehadiran

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/usecase"
	"gorm.io/gorm"
)

func ProvideKehadiran(app *fiber.App, db *gorm.DB, validator *config.Validator) {
	kehadiranRepository := postgres.NewKehadiranRepository(db)
	kehadiranUseCase := usecase.NewKehadiranUseCase(&kehadiranRepository)
	kehadiranController := controller.NewKehadiranController(kehadiranUseCase, validator)

	jadwalRepository := postgres.NewJadwalRepository(db)
	jadwalUseCase := usecase.NewJadwalUseCase(&jadwalRepository)
	jadwalController := controller.NewJadwalController(jadwalUseCase, validator)

	router.Route(
		app,
		kehadiranController,
		jadwalController,
	)
}