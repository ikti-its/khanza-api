package controller

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type JadwalController struct {
	UseCase   *usecase.JadwalUseCase
	Validator *config.Validator
}

func NewJadwalController(useCase *usecase.JadwalUseCase, validator *config.Validator) *JadwalController {
	return &JadwalController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *JadwalController) Get(ctx *fiber.Ctx) error {
	jadwal := c.UseCase.Get()

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   jadwal,
	})
}

func (c *JadwalController) GetByHariId(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request parameter",
		})
	}

	jadwal := c.UseCase.GetByHariId(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   jadwal,
	})
}

func (c *JadwalController) GetByPegawaiId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	jadwal := c.UseCase.GetByPegawaiId(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   jadwal,
	})
}

func (c *JadwalController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	jadwal := c.UseCase.GetById(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   jadwal,
	})
}

func (c *JadwalController) Update(ctx *fiber.Ctx) error {
	var request model.UpdateJadwalRequest
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

	updater := ctx.Locals("user").(string)
	response := c.UseCase.Update(&request, id, updater)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
