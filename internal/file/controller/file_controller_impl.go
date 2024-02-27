package controller

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/fathoor/simkes-api/internal/file/model"
	"github.com/fathoor/simkes-api/internal/file/service"
	"github.com/gofiber/fiber/v2"
)

type fileControllerImpl struct {
	service.FileService
}

func (controller *fileControllerImpl) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "No file uploaded",
		})
	}

	fileType := c.FormValue("type")

	request := model.FileRequest{
		File: file,
		Type: fileType,
	}

	response := controller.FileService.Upload(&request)

	if err := c.SaveFile(file, response.Path); err != nil {
		panic(exception.InternalServerError{
			Message: "Failed to save file",
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *fileControllerImpl) Download(c *fiber.Ctx) error {
	fileType := c.Params("filetype")
	fileName := c.Params("filename")

	filePath := controller.FileService.Get(fileType, fileName)

	return c.Download(filePath)
}

func (controller *fileControllerImpl) View(c *fiber.Ctx) error {
	fileType := c.Params("filetype")
	fileName := c.Params("filename")

	filePath := controller.FileService.Get(fileType, fileName)

	return c.SendFile(filePath)
}

func (controller *fileControllerImpl) Delete(c *fiber.Ctx) error {
	fileType := c.Params("filetype")
	fileName := c.Params("filename")

	controller.FileService.Delete(fileType, fileName)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewFileControllerProvider(service *service.FileService) FileController {
	return &fileControllerImpl{*service}
}
