package controller

import (
	"github.com/gofiber/fiber/v2"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/ref/internal/usecase"
)

type RefController struct {
	UseCase *usecase.RefUseCase
}

func NewRefController(useCase *usecase.RefUseCase) *RefController {
	return &RefController{
		UseCase: useCase,
	}
}

func (c *RefController) GetRole(ctx *fiber.Ctx) error {
	response := c.UseCase.GetRole()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetJabatan(ctx *fiber.Ctx) error {
	response := c.UseCase.GetJabatan()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetDepartemen(ctx *fiber.Ctx) error {
	response := c.UseCase.GetDepartemen()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetStatusAktif(ctx *fiber.Ctx) error {
	response := c.UseCase.GetStatusAktif()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetShift(ctx *fiber.Ctx) error {
	response := c.UseCase.GetShift()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetAlasanCuti(ctx *fiber.Ctx) error {
	response := c.UseCase.GetAlasanCuti()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
