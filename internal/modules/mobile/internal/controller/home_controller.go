package controller

import (
	"github.com/gofiber/fiber/v2"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/usecase"
)

type HomeController struct {
	UseCase *usecase.HomeUseCase
}

func NewHomeController(useCase *usecase.HomeUseCase) *HomeController {
	return &HomeController{
		UseCase: useCase,
	}
}

func (c *HomeController) GetHomePegawai(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	hari := ctx.QueryInt("hari", 0)
	tanggal := ctx.Query("tanggal")

	response := c.UseCase.GetHomePegawai(id, hari, tanggal)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
