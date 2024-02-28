package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/role/model"
	"github.com/fathoor/simkes-api/internal/role/service"
	"github.com/gofiber/fiber/v2"
)

type roleControllerImpl struct {
	service.RoleService
}

func (controller *roleControllerImpl) Create(c *fiber.Ctx) error {
	var request model.RoleRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.RoleService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *roleControllerImpl) Get(c *fiber.Ctx) error {
	response := controller.RoleService.GetAll()

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *roleControllerImpl) GetByNama(c *fiber.Ctx) error {
	role := c.Params("role")

	response := controller.RoleService.GetByRole(role)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *roleControllerImpl) Update(c *fiber.Ctx) error {
	var request model.RoleRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	role := c.Params("role")

	response := controller.RoleService.Update(role, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *roleControllerImpl) Delete(c *fiber.Ctx) error {
	role := c.Params("role")

	controller.RoleService.Delete(role)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewRoleControllerProvider(service *service.RoleService) RoleController {
	return &roleControllerImpl{*service}
}
