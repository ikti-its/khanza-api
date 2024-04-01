package middleware

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func Authenticate(roles []int) func(ctx *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SuccessHandler: func(ctx *fiber.Ctx) error {
			claims := ctx.Locals("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

			ctx.Locals("user", claims["sub"])
			ctx.Locals("role", int(claims["role"].(float64)))

			// Roles
			// 0: Public, 1337: Developer, 1: Admin, 2: Pegawai, ... (ask developer)
			role := int(claims["role"].(float64))

			if roles[0] == 0 {
				return ctx.Next()
			}

			for _, r := range roles {
				if role == r {
					return ctx.Next()
				}
			}

			panic(&exception.ForbiddenError{
				Message: "You don't have permission to access this resource",
			})

			return nil
		},

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				panic(&exception.UnauthorizedError{
					Message: "Missing or malformed JWT",
				})
			} else {
				panic(&exception.UnauthorizedError{
					Message: "Invalid or expired JWT",
				})
			}
		},

		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS512.Alg(),
			Key:    []byte(secret),
		},

		ContextKey: "jwt",
	})
}
