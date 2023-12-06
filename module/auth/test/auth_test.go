package test

import (
	"encoding/json"
	"github.com/fathoor/simkes-api/core/config"
	"github.com/fathoor/simkes-api/core/provider"
	"github.com/fathoor/simkes-api/module/auth/model"
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

	provider.ProvideAuth(app, db)

	return app
}

var app = ProvideTestApp()

func TestAuth_Login(t *testing.T) {
	t.Run("When request is valid", func(t *testing.T) {
		authRequest := model.AuthRequest{
			NIP:      "Admin",
			Password: "admin",
		}
		requestBody, err := json.Marshal(authRequest)
		assert.Nil(t, err)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/auth", strings.NewReader(string(requestBody)))
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

		request := httptest.NewRequest(http.MethodPost, "/api/v1/auth", strings.NewReader(string(requestBody)))
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

		request := httptest.NewRequest(http.MethodPost, "/api/v1/auth", strings.NewReader(string(requestBody)))
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

		request := httptest.NewRequest(http.MethodPost, "/api/v1/auth", strings.NewReader(string(requestBody)))
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, response.StatusCode)
	})
}
