package controller

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/middleware"
	web "github.com/fathoor/simkes-api/core/model"
	"github.com/fathoor/simkes-api/module/file/model"
	"github.com/fathoor/simkes-api/module/file/service"
	"github.com/gofiber/fiber/v2"
)

type fileControllerImpl struct {
	service.FileService
}

func (controller *fileControllerImpl) Route(app *fiber.App) {
	file := app.Group("/api/v1/file", middleware.Authenticate(0))

	file.Post("/", controller.Upload)
	file.Get("/:filetype/:filename/download", controller.Download)
	file.Get("/:filetype/:filename", controller.View)
	file.Delete("/:filetype/:filename", controller.Delete)
}

func (controller *fileControllerImpl) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	exception.PanicIfError(err)

	fileType := c.FormValue("type")

	request := model.FileRequest{
		File: file,
		Type: fileType,
	}

	response := controller.FileService.Upload(&request)

	err = c.SaveFile(file, response.URL)
	exception.PanicIfError(err)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *fileControllerImpl) Download(c *fiber.Ctx) error {
	fileType := c.Params("filetype")
	fileName := c.Params("filename")

	filePath, err := controller.FileService.Get(fileType, fileName)
	exception.PanicIfError(err)

	return c.Download(filePath)
}

func (controller *fileControllerImpl) View(c *fiber.Ctx) error {
	fileType := c.Params("filetype")
	fileName := c.Params("filename")

	filePath, err := controller.FileService.Get(fileType, fileName)
	exception.PanicIfError(err)

	return c.SendFile(filePath)
}

func (controller *fileControllerImpl) Delete(c *fiber.Ctx) error {
	fileType := c.Params("filetype")
	fileName := c.Params("filename")

	err := controller.FileService.Delete(fileType, fileName)
	exception.PanicIfError(err)

	return c.SendStatus(fiber.StatusNoContent)
}

func ProvideFileController(service *service.FileService) FileController {
	return &fileControllerImpl{*service}
}
