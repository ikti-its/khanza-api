package controller

import (
	"github.com/fathoor/simkes-api/core/exception"
	web "github.com/fathoor/simkes-api/core/model"
	"github.com/fathoor/simkes-api/module/auth/model"
	"github.com/fathoor/simkes-api/module/auth/service"
	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct {
	service.AuthService
}

func (controller *authControllerImpl) Route(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/", controller.Login)
}

func (controller *authControllerImpl) Login(c *fiber.Ctx) error {
	var request model.AuthRequest

	parse := c.BodyParser(&request)
	if parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response, _ := controller.AuthService.Login(&request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func ProvideAuthController(service *service.AuthService) AuthController {
	return &authControllerImpl{*service}
}
