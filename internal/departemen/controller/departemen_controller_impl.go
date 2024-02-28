package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/departemen/model"
	"github.com/fathoor/simkes-api/internal/departemen/service"
	"github.com/gofiber/fiber/v2"
)

type departemenControllerImpl struct {
	service.DepartemenService
}

func (controller *departemenControllerImpl) Create(c *fiber.Ctx) error {
	var request model.DepartemenRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.DepartemenService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *departemenControllerImpl) Get(c *fiber.Ctx) error {
	response := controller.DepartemenService.GetAll()

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *departemenControllerImpl) GetByNama(c *fiber.Ctx) error {
	departemen := c.Params("departemen")

	response := controller.DepartemenService.GetByDepartemen(departemen)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *departemenControllerImpl) Update(c *fiber.Ctx) error {
	var request model.DepartemenRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	departemen := c.Params("departemen")

	response := controller.DepartemenService.Update(departemen, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *departemenControllerImpl) Delete(c *fiber.Ctx) error {
	departemen := c.Params("departemen")

	controller.DepartemenService.Delete(departemen)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewDepartemenControllerProvider(service *service.DepartemenService) DepartemenController {
	return &departemenControllerImpl{*service}
}
