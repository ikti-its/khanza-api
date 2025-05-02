package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/stokobatpasien/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/stokobatpasien/internal/usecase"
)

type StokObatPasienController struct {
	UseCase *usecase.StokObatPasienUseCase
}

func NewStokObatPasienController(useCase *usecase.StokObatPasienUseCase) *StokObatPasienController {
	return &StokObatPasienController{
		UseCase: useCase,
	}
}

func (c *StokObatPasienController) Create(ctx *fiber.Ctx) error {
	var request model.StokObatPasienRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid request body",
		})
	}

	response, err := c.UseCase.Create(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *StokObatPasienController) GetAll(ctx *fiber.Ctx) error {
	result, err := c.UseCase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (c *StokObatPasienController) GetByNoPermintaan(ctx *fiber.Ctx) error {
	noPermintaan := ctx.Params("no_permintaan")

	result, err := c.UseCase.GetByNoPermintaan(noPermintaan)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   fmt.Sprintf("no_permintaan not found: %v", err),
		})
	}
	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (c *StokObatPasienController) Update(ctx *fiber.Ctx) error {
	var request model.StokObatPasienRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid request",
		})
	}

	response, err := c.UseCase.Update(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *StokObatPasienController) Delete(ctx *fiber.Ctx) error {
	noPermintaan := ctx.Params("no_permintaan")

	if noPermintaan == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "no_permintaan is required",
		})
	}

	err := c.UseCase.DeleteByNoPermintaan(noPermintaan)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Stok obat pasien deleted successfully",
	})
}

func (c *StokObatPasienController) GetByNomorRawat(ctx *fiber.Ctx) error {
	nomorRawat := ctx.Params("nomor_rawat")

	data, err := c.UseCase.GetByNomorRawat(nomorRawat)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   fmt.Sprintf("no_rawat not found: %v", err),
		})
	}

	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}
