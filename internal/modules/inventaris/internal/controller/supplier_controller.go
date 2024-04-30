package controller

import (
	"github.com/gofiber/fiber/v2"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/usecase"
)

type SupplierController struct {
	UseCase *usecase.SupplierUseCase
}

func NewSupplierController(useCase *usecase.SupplierUseCase) *SupplierController {
	return &SupplierController{
		UseCase: useCase,
	}
}

func (c *SupplierController) Get(ctx *fiber.Ctx) error {
	response := c.UseCase.Get()

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
