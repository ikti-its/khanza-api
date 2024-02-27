package middleware

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeNIP() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		nip := claims["nip"].(string)

		if role == "Admin" {
			return c.Next()
		}

		if nip == c.Params("nip") {
			return c.Next()
		} else {
			panic(exception.UnauthorizedError{
				Message: "Unauthorized",
			})
		}
	}
}
