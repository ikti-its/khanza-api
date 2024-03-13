package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/cuti/model"
	"github.com/fathoor/simkes-api/internal/cuti/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type cutiControllerImpl struct {
	service.CutiService
}

func (controller *cutiControllerImpl) Create(c *fiber.Ctx) error {
	var request model.CutiCreateRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.CutiService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *cutiControllerImpl) Get(c *fiber.Ctx) error {
	nip := c.Query("nip")

	if nip != "" {
		response := controller.CutiService.GetByNIP(nip)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := controller.CutiService.GetAll()

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *cutiControllerImpl) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	response := controller.CutiService.GetByID(id)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *cutiControllerImpl) Update(c *fiber.Ctx) error {
	var request model.CutiUpdateRequest

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	id := c.Params("id")

	if role == "Admin" {
		response := controller.CutiService.UpdateStatus(id, &request)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := controller.CutiService.Update(id, &request)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *cutiControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.CutiService.Delete(id)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewCutiControllerProvider(service *service.CutiService) CutiController {
	return &cutiControllerImpl{*service}
}
