package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/shift/model"
	"github.com/fathoor/simkes-api/internal/shift/service"
	"github.com/gofiber/fiber/v2"
)

type shiftControllerImpl struct {
	service.ShiftService
}

func (controller *shiftControllerImpl) Create(c *fiber.Ctx) error {
	var request model.ShiftRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.ShiftService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) Get(c *fiber.Ctx) error {
	response := controller.ShiftService.GetAll()

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) GetByNama(c *fiber.Ctx) error {
	shift := c.Params("shift")

	response := controller.ShiftService.GetByNama(shift)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) Update(c *fiber.Ctx) error {
	var request model.ShiftRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	shift := c.Params("shift")

	response := controller.ShiftService.Update(shift, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) Delete(c *fiber.Ctx) error {
	shift := c.Params("shift")

	controller.ShiftService.Delete(shift)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewShiftControllerProvider(service *service.ShiftService) ShiftController {
	return &shiftControllerImpl{*service}
}
