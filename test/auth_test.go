package test

import (
	"encoding/json"
	"github.com/fathoor/simkes-api/internal/auth/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuth_Login(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		authRequest := model.AuthRequest{
			NIP:      "Admin",
			Password: "admin",
		}
		requestBody, err := json.Marshal(authRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/auth", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
	})

	t.Run("When request is invalid", func(t *testing.T) {
		authRequest := model.AuthRequest{
			NIP:      "",
			Password: "",
		}
		requestBody, err := json.Marshal(authRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/auth", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	})

	t.Run("When NIP is invalid", func(t *testing.T) {
		authRequest := model.AuthRequest{
			NIP:      "D0001",
			Password: "admin",
		}
		requestBody, err := json.Marshal(authRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/auth", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	})

	t.Run("When password is invalid", func(t *testing.T) {
		authRequest := model.AuthRequest{
			NIP:      "Admin",
			Password: "admin123",
		}
		requestBody, err := json.Marshal(authRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/v1/auth", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, response.StatusCode)
	})
}
