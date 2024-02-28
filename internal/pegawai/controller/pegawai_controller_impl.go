package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/pegawai/model"
	"github.com/fathoor/simkes-api/internal/pegawai/service"
	"github.com/gofiber/fiber/v2"
)

type pegawaiControllerImpl struct {
	service.PegawaiService
}

func (controller *pegawaiControllerImpl) Create(c *fiber.Ctx) error {
	var request model.PegawaiRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.PegawaiService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *pegawaiControllerImpl) Get(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	size := c.QueryInt("size")

	if size < 10 {
		size = 10
	}

	if page < 1 {
		response := controller.PegawaiService.GetAll()

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := controller.PegawaiService.GetPage(page, size)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *pegawaiControllerImpl) GetByNIP(c *fiber.Ctx) error {
	nip := c.Params("nip")

	response := controller.PegawaiService.GetByNIP(nip)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pegawaiControllerImpl) Update(c *fiber.Ctx) error {
	var request model.PegawaiRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := c.Params("nip")

	response := controller.PegawaiService.Update(nip, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pegawaiControllerImpl) Delete(c *fiber.Ctx) error {
	nip := c.Params("nip")

	controller.PegawaiService.Delete(nip)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewPegawaiControllerProvider(service *service.PegawaiService) PegawaiController {
	return &pegawaiControllerImpl{*service}
}
