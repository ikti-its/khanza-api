package test

import (
	"errors"
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestException_Panic(t *testing.T) {
	t.Run("When error then panic", func(t *testing.T) {
		assert.Panics(t, func() {
			exception.PanicIfError(errors.New("test"))
		})
	})

	t.Run("When no error then nothing", func(t *testing.T) {
		assert.NotPanics(t, func() {
			exception.PanicIfError(nil)
		})
	})
}

func TestException_Handler(t *testing.T) {
	t.Run("When exception error then handle recover", func(t *testing.T) {
		app.Get("/", func(c *fiber.Ctx) error {
			return exception.BadRequestError{
				Message: "Bad Request Error",
			}
		})

		request := httptest.NewRequest("GET", "/", nil)
		response, err := app.Test(request)
		assert.Nil(t, err)

		bytes, err := io.ReadAll(response.Body)
		assert.Nil(t, err)
		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, "{\"code\":400,\"data\":\"Bad Request Error\",\"message\":\"Bad Request\"}", string(bytes))
	})
}
