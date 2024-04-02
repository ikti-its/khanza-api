package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/usecase"
	"gorm.io/gorm"
)

func ProvideAuth(app *fiber.App, cfg *config.Config, db *gorm.DB, validator *config.Validator) {
	authRepository := postgres.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(&authRepository, cfg)
	authController := controller.NewAuthController(authUseCase, validator)

	router.Route(app, authController)
}
