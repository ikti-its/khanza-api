package controller

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/middleware"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type FotoController struct {
	UseCase   *usecase.FotoUseCase
	Validator *config.Validator
}

func NewFotoController(useCase *usecase.FotoUseCase, validator *config.Validator) *FotoController {
	return &FotoController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *FotoController) Create(ctx *fiber.Ctx) error {
	var request model.FotoRequest

	if err := ctx.BodyParser(&request); err != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	akunId := c.UseCase.GetAkunId(request.IdPegawai)
	middleware.AuthorizeUserFoto(akunId)

	if err := c.Validator.Validate(&request); err != nil {
		message := c.Validator.Message(err)
		panic(&exception.BadRequestError{
			Message: message,
		})
	}

	updater := ctx.Locals("user").(string)
	response := c.UseCase.Create(&request, updater)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *FotoController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	middleware.AuthorizeUserFoto(id)

	response := c.UseCase.GetById(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *FotoController) Update(ctx *fiber.Ctx) error {
	var request model.FotoRequest
	id := ctx.Params("id")

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

	middleware.AuthorizeUserFoto(id)

	updater := ctx.Locals("user").(string)
	response := c.UseCase.Update(&request, id, updater)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *FotoController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.UseCase.Delete(id)

	return ctx.SendStatus(fiber.StatusNoContent)
}
