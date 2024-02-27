package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/auth/model"
	"github.com/fathoor/simkes-api/internal/auth/service"
	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct {
	service.AuthService
}

func (controller *authControllerImpl) Login(c *fiber.Ctx) error {
	var request model.AuthRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.AuthService.Login(&request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func NewAuthControllerProvider(service *service.AuthService) AuthController {
	return &authControllerImpl{*service}
}
