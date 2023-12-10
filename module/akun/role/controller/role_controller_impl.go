package controller

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/middleware"
	web "github.com/fathoor/simkes-api/core/model"
	"github.com/fathoor/simkes-api/module/akun/role/model"
	"github.com/fathoor/simkes-api/module/akun/role/service"
	"github.com/gofiber/fiber/v2"
)

type roleControllerImpl struct {
	service.RoleService
}

func (controller *roleControllerImpl) Route(app *fiber.App) {
	role := app.Group("/v1/akun/role", middleware.Authenticate("Admin"))

	role.Post("/", controller.Create)
	role.Get("/", controller.GetAll)
	role.Get("/:id", controller.Get)
	role.Put("/:id", controller.Update)
	role.Delete("/:id", controller.Delete)
}

func (controller *roleControllerImpl) Create(c *fiber.Ctx) error {
	var request model.RoleRequest

	parse := c.BodyParser(&request)
	if parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	err := controller.RoleService.Create(&request)
	exception.PanicIfError(err)

	return c.SendStatus(fiber.StatusCreated)
}

func (controller *roleControllerImpl) GetAll(c *fiber.Ctx) error {
	response, err := controller.RoleService.GetAll()
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *roleControllerImpl) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid id",
		})
	}

	response, _ := controller.RoleService.GetByID(id)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *roleControllerImpl) Update(c *fiber.Ctx) error {
	var request model.RoleRequest

	parse := c.BodyParser(&request)
	if parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid id",
		})
	}

	response, err := controller.RoleService.Update(id, &request)
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *roleControllerImpl) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid id",
		})
	}

	err = controller.RoleService.Delete(id)
	exception.PanicIfError(err)

	return c.SendStatus(fiber.StatusNoContent)
}

func ProvideRoleController(service *service.RoleService) RoleController {
	return &roleControllerImpl{*service}
}
