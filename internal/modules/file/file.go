package file

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/modules/file/internal/controller"
	"github.com/fathoor/simkes-api/internal/modules/file/internal/router"
	"github.com/fathoor/simkes-api/internal/modules/file/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func ProvideFile(app *fiber.App, cfg *config.Config, validator *config.Validator) {
	fileUseCase := usecase.NewFileUseCase(cfg)
	fileController := controller.NewFileController(fileUseCase, validator)

	router.Route(app, fileController)
}
