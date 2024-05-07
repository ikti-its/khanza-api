package controller

import (
	"github.com/gofiber/fiber/v2"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/usecase"
)

type SatuanController struct {
	UseCase *usecase.SatuanUseCase
}

func NewSatuanController(useCase *usecase.SatuanUseCase) *SatuanController {
	return &SatuanController{UseCase: useCase}
}

func (c *SatuanController) Get(ctx *fiber.Ctx) error {
	response := c.UseCase.Get()

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
