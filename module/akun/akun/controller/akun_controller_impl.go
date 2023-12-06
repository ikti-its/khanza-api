package controller

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/middleware"
	web "github.com/fathoor/simkes-api/core/model"
	"github.com/fathoor/simkes-api/module/akun/akun/model"
	"github.com/fathoor/simkes-api/module/akun/akun/service"
	"github.com/gofiber/fiber/v2"
)

type akunControllerImpl struct {
	service.AkunService
}

func (controller *akunControllerImpl) Route(app *fiber.App) {
	akun := app.Group("/api/v1/akun", middleware.Authenticate(1))

	akun.Post("/", controller.Create)
	akun.Get("/", controller.GetAll)
	akun.Get("/:nip", controller.Get)
	akun.Put("/:nip", controller.Update)
	akun.Delete("/:nip", controller.Delete)
}

func (controller *akunControllerImpl) Create(c *fiber.Ctx) error {
	var request model.AkunRequest

	parse := c.BodyParser(&request)
	exception.PanicIfError(parse)

	err := controller.AkunService.Create(&request)
	exception.PanicIfError(err)

	return c.SendStatus(fiber.StatusCreated)
}

func (controller *akunControllerImpl) GetAll(c *fiber.Ctx) error {
	response, err := controller.AkunService.GetAll()
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) Get(c *fiber.Ctx) error {
	nip := c.Params("nip")

	response, err := controller.AkunService.GetByNIP(nip)
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) Update(c *fiber.Ctx) error {
	var request model.AkunRequest

	parse := c.BodyParser(&request)
	exception.PanicIfError(parse)

	nip := c.Params("nip")

	response, err := controller.AkunService.Update(nip, &request)
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) Delete(c *fiber.Ctx) error {
	nip := c.Params("nip")

	err := controller.AkunService.Delete(nip)
	exception.PanicIfError(err)

	return c.SendStatus(fiber.StatusNoContent)
}

func ProvideAkunController(service *service.AkunService) AkunController {
	return &akunControllerImpl{*service}
}
