package exception

import (
	"errors"
	web "github.com/fathoor/simkes-api/internal/app/model"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, e error) error {
	var (
		badRequestError     BadRequestError
		unauthorizedError   UnauthorizedError
		forbiddenError      ForbiddenError
		notFoundError       NotFoundError
		internalServerError InternalServerError
	)

	switch {
	case errors.As(e, &badRequestError):
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   badRequestError.Error(),
		})
	case errors.As(e, &unauthorizedError):
		return c.Status(fiber.StatusUnauthorized).JSON(web.Response{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   unauthorizedError.Error(),
		})
	case errors.As(e, &forbiddenError):
		return c.Status(fiber.StatusForbidden).JSON(web.Response{
			Code:   fiber.StatusForbidden,
			Status: "Forbidden",
			Data:   forbiddenError.Error(),
		})
	case errors.As(e, &notFoundError):
		return c.Status(fiber.StatusNotFound).JSON(web.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   notFoundError.Error(),
		})
	case errors.As(e, &internalServerError):
		return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   internalServerError.Error(),
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   e.Error(),
		})
	}
}
