package test

import (
	"bytes"
	_ "embed"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed assets/default.png
var fileTest []byte

//go:embed assets/oversize.png
var fileOversizeTest []byte

func TestFile_Upload(t *testing.T) {
	t.Run("When file and type is valid", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		file, err := writer.CreateFormFile("file", "default.png")
		assert.Nil(t, err)
		_, err = file.Write(fileTest)
		assert.Nil(t, err)

		err = writer.WriteField("type", "image")
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/file", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When file is invalid", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		file, err := writer.CreateFormFile("file", "default.bmp")
		assert.Nil(t, err)
		_, err = file.Write(fileTest)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/file", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When type is invalid", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		file, err := writer.CreateFormFile("file", "default.png")
		assert.Nil(t, err)
		_, err = file.Write(fileTest)
		assert.Nil(t, err)

		err = writer.WriteField("type", "video")
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/file", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When file size exceeds limit", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		file, err := writer.CreateFormFile("file", "oversize.png")
		assert.Nil(t, err)
		_, err = file.Write(fileOversizeTest)
		assert.Nil(t, err)

		err = writer.WriteField("type", "image")
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/file", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})
}

func TestFile_Download(t *testing.T) {
	t.Run("When file exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/file/image/default.png/download", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When file not exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/file/image/default.bmp/download", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestFile_Get(t *testing.T) {
	t.Run("When file exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/file/image/default.png", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When file not exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/file/image/default.bmp", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestFile_Delete(t *testing.T) {
	t.Run("When file exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/file/image/default.png", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
	})

	t.Run("When file not exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/file/image/default.bmp", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}
