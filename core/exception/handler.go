package exception

import (
	web "github.com/fathoor/simkes-api/core/model"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, e error) error {
	switch e.(type) {
	case BadRequestError:
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   e.Error(),
		})
	case UnauthorizedError:
		return c.Status(fiber.StatusUnauthorized).JSON(web.Response{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   e.Error(),
		})
	case ForbiddenError:
		return c.Status(fiber.StatusForbidden).JSON(web.Response{
			Code:   fiber.StatusForbidden,
			Status: "Forbidden",
			Data:   e.Error(),
		})
	case NotFoundError:
		return c.Status(fiber.StatusNotFound).JSON(web.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   e.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
		Code:   fiber.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   e.Error(),
	})
}
