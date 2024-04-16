package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/app/exception"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/usecase"
)

type MedisController struct {
	UseCase   *usecase.MedisUseCase
	Validator *config.Validator
}

func NewMedisController(useCase *usecase.MedisUseCase, validator *config.Validator) *MedisController {
	return &MedisController{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (c *MedisController) Create(ctx *fiber.Ctx) error {
	var request model.MedisRequest

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
	response := c.UseCase.Create(&request, updater)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *MedisController) Get(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page")
	size := ctx.QueryInt("size")
	jenis := ctx.Query("jenis")

	if size < 10 {
		size = 10
	}

	if page < 1 {
		response := c.UseCase.Get()

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		if jenis != "" {
			response := c.UseCase.GetByJenis(jenis)

			return ctx.Status(fiber.StatusOK).JSON(web.Response{
				Code:   fiber.StatusOK,
				Status: "OK",
				Data:   response,
			})
		} else {
			response := c.UseCase.GetPage(page, size)

			return ctx.Status(fiber.StatusOK).JSON(web.Response{
				Code:   fiber.StatusOK,
				Status: "OK",
				Data:   response,
			})
		}
	}
}

func (c *MedisController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.UseCase.GetById(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *MedisController) Update(ctx *fiber.Ctx) error {
	var request model.MedisRequest
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

func (c *MedisController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	updater := ctx.Locals("user").(string)

	c.UseCase.Delete(id, updater)

	return ctx.SendStatus(fiber.StatusNoContent)
}
