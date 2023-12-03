package exception

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx, e error) error {
	switch e.(type) {
	case BadRequestError:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "Bad Request",
			"data":    e.Error(),
		})
	case UnauthorizedError:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "Unauthorized",
			"data":    e.Error(),
		})
	case ForbiddenError:
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    403,
			"message": "Forbidden",
			"data":    e.Error(),
		})
	case NotFoundError:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    404,
			"message": "Not Found",
			"data":    e.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    500,
		"message": "Internal Server Error",
		"data":    e.Error(),
	})
}
