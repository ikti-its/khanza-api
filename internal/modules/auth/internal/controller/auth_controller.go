package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/app/exception"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/auth/internal/usecase"
)

type AuthController struct {
	UseCase   *usecase.AuthUseCase
	Validator *config.Validator
}

func NewAuthController(useCase *usecase.AuthUseCase, validator *config.Validator) *AuthController {
	return &AuthController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var request model.AuthRequest

	if err := ctx.BodyParser(&request); err != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	if err := c.Validator.Validate(&request); err != nil {
		message := c.Validator.Message(err)
		panic(&exception.BadRequestError{
			Message: message,
		})
	}

	response := c.UseCase.Login(&request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
