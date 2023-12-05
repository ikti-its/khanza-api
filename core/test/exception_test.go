package test

import (
	"encoding/json"
	"errors"
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
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

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response, err := app.Test(request)
		assert.Nil(t, err)

		var webResponse model.Response
		err = json.NewDecoder(response.Body).Decode(&webResponse)
		assert.Nil(t, err)
		assert.Equal(t, model.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Bad Request Error",
		}, webResponse)
	})
}
