package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/exception"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/usecase"
)

type RegistrasiController struct {
	UseCase *usecase.RegistrasiUseCase
}

func NewRegistrasiController(useCase *usecase.RegistrasiUseCase) *RegistrasiController {
	return &RegistrasiController{
		UseCase: useCase,
	}
}

func (c *RegistrasiController) Create(ctx *fiber.Ctx) error {
	var request model.RegistrasiRequest
	fmt.Println("Received a POST request to /registrasi") // Debugging log

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("Failed to parse request body:", err) // Debugging log
		if err := ctx.BodyParser(&request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
				Code:   fiber.StatusBadRequest,
				Status: "Bad Request",
				Data:   "Invalid request body",
			})
		}

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
		Status: "OK",
		Data:   response,
	})
}

func (c *RegistrasiController) GetAll(ctx *fiber.Ctx) error {
	response, err := c.UseCase.GetAll()
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

func (c *RegistrasiController) GetByNomorReg(ctx *fiber.Ctx) error {
	nomorReg := ctx.Params("nomor_reg")
	response, err := c.UseCase.GetByNomorReg(nomorReg)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *RegistrasiController) Update(ctx *fiber.Ctx) error {
	nomorReg := ctx.Params("nomor_reg")
	var request model.RegistrasiRequest

	if err := ctx.BodyParser(&request); err != nil {
		panic(&exception.BadRequestError{Message: "Invalid request body"})
	}

	response, err := c.UseCase.Update(nomorReg, &request)
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

func (c *RegistrasiController) Delete(ctx *fiber.Ctx) error {
	nomorReg := ctx.Params("nomor_reg")
	err := c.UseCase.Delete(nomorReg)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(web.Response{
		Code:   fiber.StatusNoContent,
		Status: "Deleted",
	})
}
