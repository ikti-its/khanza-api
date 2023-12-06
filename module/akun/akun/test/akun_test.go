package test

import (
	"encoding/json"
	"github.com/fathoor/simkes-api/core/config"
	"github.com/fathoor/simkes-api/core/provider"
	"github.com/fathoor/simkes-api/module/akun/akun/model"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ProvideTestApp() *fiber.App {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp(cfg)
		db  = config.ProvideDB(cfg)
	)

	provider.ProvideAkun(app, db)

	return app
}

var app = ProvideTestApp()

func TestRole_Create(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "Admin",
			Email:    "admin@fathoor.cloud",
			Password: "admin",
			RoleID:   1,
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/akun", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "Admin",
			Email:    "",
			Password: "admin",
			RoleID:   1,
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/akun", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})
}

func TestRole_GetAll(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/akun", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestRole_Get(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/v1/akun/Admin", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/v1/akun/123", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestRole_Update(t *testing.T) {
	t.Run("When request and ID is valid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "Admin",
			Email:    "admin@fathoor.cloud",
			Password: "admin",
			RoleID:   1,
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/api/v1/akun/Admin", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "Admin",
			Email:    "",
			Password: "admin",
			RoleID:   1,
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/api/v1/akun/Admin", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "Admin",
			Email:    "admin@fathoor.cloud",
			Password: "admin",
			RoleID:   1,
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/api/v1/akun/123", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestRole_Delete(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/akun/Admin", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/akun/123", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}
