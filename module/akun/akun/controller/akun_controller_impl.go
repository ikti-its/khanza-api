package controller

import (
	"github.com/fathoor/simkes-api/core/exception"
	web "github.com/fathoor/simkes-api/core/model"
	"github.com/fathoor/simkes-api/module/akun/akun/model"
	"github.com/fathoor/simkes-api/module/akun/akun/service"
	"github.com/gofiber/fiber/v2"
)

type akunControllerImpl struct {
	service.AkunService
}

func (controller *akunControllerImpl) Route(app *fiber.App) {
	akun := app.Group("/v1/akun")

	akun.Post("/", controller.Create)
	akun.Get("/", controller.GetAll)
	akun.Get("/detail/:nip", controller.Get)
	akun.Put("/detail/:nip", controller.Update)
	akun.Delete("/detail/:nip", controller.Delete)

	pegawai := app.Group("/v1/akun/pegawai")

	pegawai.Get("/detail/:nip", controller.PegawaiGet)
	pegawai.Put("/detail/:nip", controller.PegawaiUpdate)
}

func (controller *akunControllerImpl) Create(c *fiber.Ctx) error {
	var request model.AkunRequest

	parse := c.BodyParser(&request)
	if parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

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

	response, _ := controller.AkunService.GetByNIP(nip)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) PegawaiGet(c *fiber.Ctx) error {
	nip := c.Params("nip")

	response, _ := controller.AkunService.PegawaiGetByNIP(nip)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) Update(c *fiber.Ctx) error {
	var request model.AkunRequest

	parse := c.BodyParser(&request)
	if parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := c.Params("nip")

	response, err := controller.AkunService.Update(nip, &request)
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) PegawaiUpdate(c *fiber.Ctx) error {
	var request model.AkunUpdateRequest

	parse := c.BodyParser(&request)
	if parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := c.Params("nip")

	response, err := controller.AkunService.PegawaiUpdate(nip, &request)
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
