package controller

import (
	"github.com/fathoor/simkes-api/internal/akun/model"
	"github.com/fathoor/simkes-api/internal/akun/service"
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type akunControllerImpl struct {
	service.AkunService
}

func (controller *akunControllerImpl) Create(c *fiber.Ctx) error {
	var request model.AkunRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.AkunService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *akunControllerImpl) Get(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	size := c.QueryInt("size")

	if size < 10 {
		size = 10
	}

	if page < 1 {
		response := controller.AkunService.GetAll()

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := controller.AkunService.GetPage(page, size)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *akunControllerImpl) GetByNIP(c *fiber.Ctx) error {
	nip := c.Params("nip")

	response := controller.AkunService.GetByNIP(nip)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *akunControllerImpl) Update(c *fiber.Ctx) error {
	var request model.AkunRequest

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := c.Params("nip")

	if role == "Admin" {
		response := controller.AkunService.UpdateAdmin(nip, &request)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := controller.AkunService.Update(nip, &request)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *akunControllerImpl) Delete(c *fiber.Ctx) error {
	nip := c.Params("nip")

	controller.AkunService.Delete(nip)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewAkunControllerProvider(service *service.AkunService) AkunController {
	return &akunControllerImpl{*service}
}
