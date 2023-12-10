package test

import (
	"encoding/json"
	"github.com/fathoor/simkes-api/core/config"
	"github.com/fathoor/simkes-api/core/provider"
	"github.com/fathoor/simkes-api/module/akun/role/model"
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

	provider.ProvideRole(app, db)

	return app
}

var app = ProvideTestApp()

func TestRole_Create(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Perawat",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun/role", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun/role", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When request is forbidden", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Admin",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun/role", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusForbidden, response.StatusCode)
	})

	t.Run("When request is duplicate", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Dokter",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun/role", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)
	})
}

func TestRole_GetAll(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/v1/akun/role", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestRole_GetByID(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/role/1", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/role/0", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestRole_Update(t *testing.T) {
	t.Run("When request and ID is valid", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Dokter",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/role/2", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/role/1", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Dokter",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/role/0", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})

	t.Run("When request is forbidden", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Admin",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/role/1", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusForbidden, response.StatusCode)
	})

	t.Run("When request is duplicate", func(t *testing.T) {
		roleRequest := model.RoleRequest{
			Name: "Dokter",
		}
		requestBody, err := json.Marshal(roleRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/role/1", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)
	})
}

func TestRole_Delete(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/akun/role/3", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/akun/role/0", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}
