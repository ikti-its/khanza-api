package test

import (
	"encoding/json"
	"github.com/fathoor/simkes-api/internal/role/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRole_Create(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "Test",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPost, "/v1/role", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		i := strings.Repeat("A", 21)
		request := model.RoleRequest{
			Nama: i,
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPost, "/v1/role", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When role already exists", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "Test",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPost, "/v1/role", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When role is Admin", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "Admin",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPost, "/v1/role", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusForbidden, response.StatusCode)
	})
}

func TestRole_GetAll(t *testing.T) {
	payload := httptest.NewRequest(http.MethodGet, "/v1/role", nil)

	response, err := app.Test(payload)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestRole_Get(t *testing.T) {
	t.Run("When role exists", func(t *testing.T) {
		payload := httptest.NewRequest(http.MethodGet, "/v1/role/Test", nil)

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When role does not exist", func(t *testing.T) {
		payload := httptest.NewRequest(http.MethodGet, "/v1/role/BEBEK", nil)

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestRole_Update(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "Test",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPut, "/v1/role/Test", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		i := strings.Repeat("A", 21)
		request := model.RoleRequest{
			Nama: i,
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPut, "/v1/role/Test", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When role does not exist", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "BEBEK",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPut, "/v1/role/BEBEK", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})

	t.Run("When role is duplicate", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "Dokter",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPut, "/v1/role/Test", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When role is Admin", func(t *testing.T) {
		request := model.RoleRequest{
			Nama: "Test",
		}

		body, err := json.Marshal(request)
		assert.Nil(t, err)

		payload := httptest.NewRequest(http.MethodPut, "/v1/role/Admin", strings.NewReader(string(body)))
		payload.Header.Set("Content-Type", "application/json")

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusForbidden, response.StatusCode)
	})
}

func TestRole_Delete(t *testing.T) {
	t.Run("When role exists", func(t *testing.T) {
		payload := httptest.NewRequest(http.MethodDelete, "/v1/role/Test", nil)

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
	})

	t.Run("When role does not exist", func(t *testing.T) {
		payload := httptest.NewRequest(http.MethodDelete, "/v1/role/BEBEK", nil)

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})

	t.Run("When role is Admin", func(t *testing.T) {
		payload := httptest.NewRequest(http.MethodDelete, "/v1/role/Admin", nil)

		response, err := app.Test(payload)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusForbidden, response.StatusCode)
	})
}
