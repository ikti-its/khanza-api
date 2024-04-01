package auth

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/controller"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/repository/postgres"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/router"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProvideAuth(app *fiber.App, cfg *config.Config, db *gorm.DB, validator *config.Validator) {
	authRepository := postgres.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(&authRepository, cfg)
	authController := controller.NewAuthController(authUseCase, validator)

	router.Route(app, authController)
}
