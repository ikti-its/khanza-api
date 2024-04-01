package controller

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/middleware"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AlamatController struct {
	UseCase   *usecase.AlamatUseCase
	Validator *config.Validator
}

func NewAlamatController(useCase *usecase.AlamatUseCase, validator *config.Validator) *AlamatController {
	return &AlamatController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *AlamatController) Create(ctx *fiber.Ctx) error {
	var request model.AlamatRequest

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

	middleware.AuthorizeUserAlamat(request.IdAkun)

	updater := ctx.Locals("user").(string)
	response := c.UseCase.Create(&request, updater)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *AlamatController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	middleware.AuthorizeUserAlamat(id)

	response := c.UseCase.GetById(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AlamatController) Update(ctx *fiber.Ctx) error {
	var request model.AlamatRequest
	id := ctx.Params("id")

	middleware.AuthorizeUserAlamat(id)

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

	updater := ctx.Locals("user").(string)
	response := c.UseCase.Update(&request, id, updater)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AlamatController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.UseCase.Delete(id)

	return ctx.SendStatus(fiber.StatusNoContent)
}
