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

func TestAkun_Create(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0004",
			Email:    "dokter4@fathoor.cloud",
			Password: "dokter",
			RoleName: "Dokter",
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
			NIP:      "D0001",
			Email:    "",
			Password: "dokter",
			RoleName: "Dokter",
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
			NIP:      "D0001",
			Email:    "dokter@fathoor.cloud",
			Password: "dokter",
			RoleName: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/akun", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)
	})
}

func TestAkun_GetAll(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/v1/akun", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestAkun_Get(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/detail/D0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/detail/P0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestAkun_PegawaiGet(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/pegawai/detail/D0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/akun/pegawai/detail/P0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestAkun_Update(t *testing.T) {
	t.Run("When request and ID is valid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0001",
			Email:    "dokter@fathoor.cloud",
			Password: "dokter",
			RoleName: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/detail/D0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0001",
			Email:    "",
			Password: "dokter",
			RoleName: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/detail/D0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		akunRequest := model.AkunRequest{
			NIP:      "D0001",
			Email:    "dokter@fathoor.cloud",
			Password: "dokter",
			RoleName: "Dokter",
		}
		requestBody, err := json.Marshal(akunRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/detail/P0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestAkun_PegawaiUpdate(t *testing.T) {
	t.Run("When request and ID is valid", func(t *testing.T) {
		akunUpdateRequest := model.AkunUpdateRequest{
			Email:    "dokter@fathoor.cloud",
			Password: "dokter",
		}
		requestBody, err := json.Marshal(akunUpdateRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/pegawai/detail/D0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		akunUpdateRequest := model.AkunUpdateRequest{
			Email:    "",
			Password: "dokter",
		}
		requestBody, err := json.Marshal(akunUpdateRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/pegawai/detail/D0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		akunUpdateRequest := model.AkunUpdateRequest{
			Email:    "pegawai@fathoor.cloud",
			Password: "pegawai",
		}
		requestBody, err := json.Marshal(akunUpdateRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPut, "/v1/akun/pegawai/detail/P0001", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}

func TestAkun_Delete(t *testing.T) {
	t.Run("When ID is valid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/akun/detail/D0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNoContent, response.StatusCode)
	})

	t.Run("When ID is not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/akun/detail/P0001", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})
}
