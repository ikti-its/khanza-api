package test

import (
	"encoding/json"
	"github.com/fathoor/simkes-api/internal/akun/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAkun_Create(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0101",
			Email:    "dokter101@fathoor.dev",
			Password: "dokter",
			RoleNama: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0101",
			Email:    "",
			Password: "dokter",
			RoleNama: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When request is duplicate", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0101",
			Email:    "dokter@fathoor.dev",
			Password: "dokter",
			RoleNama: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})
}

func TestAkun_GetAll(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When paged request is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun?page=1&size=10", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When invalid page, GetAll instead", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun?page=0&size=0", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})
}

func TestAkun_Get(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/D0101", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/P0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestAkun_Update(t *testing.T) {
	t.Run("When request and ID is valid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0101",
			Email:    "dokter@fathoor.dev",
			Password: "dokter",
			RoleNama: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/D0101", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0101",
			Email:    "",
			Password: "dokter",
			RoleNama: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/D0101", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "P0001",
			Email:    "dokter@fathoor.dev",
			Password: "dokter",
			RoleNama: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/P0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestAkun_Delete(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/akun/D0101", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/akun/P0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}
