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

type BerkasController struct {
	UseCase   *usecase.BerkasUseCase
	Validator *config.Validator
}

func NewBerkasController(useCase *usecase.BerkasUseCase, validator *config.Validator) *BerkasController {
	return &BerkasController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *BerkasController) Create(ctx *fiber.Ctx) error {
	var request model.BerkasRequest

	if err := ctx.BodyParser(&request); err != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	akunId := c.UseCase.GetAkunId(request.IdPegawai)
	middleware.AuthorizeUserBerkas(akunId)

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

func (c *BerkasController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	middleware.AuthorizeUserBerkas(id)

	response := c.UseCase.GetById(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *BerkasController) Update(ctx *fiber.Ctx) error {
	var request model.BerkasRequest
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

	middleware.AuthorizeUserBerkas(id)

	updater := ctx.Locals("user").(string)
	response := c.UseCase.Update(&request, id, updater)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *BerkasController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.UseCase.Delete(id)

	return ctx.SendStatus(fiber.StatusNoContent)
}
