package middleware

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(roles []int) func(ctx *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SuccessHandler: func(ctx *fiber.Ctx) error {
			token, ok := ctx.Locals("jwt").(*jwt.Token)
			if !ok {
				log.Println("❌ JWT token not found in context")
				return fiber.ErrUnauthorized
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("❌ JWT claims are invalid")
				return fiber.ErrUnauthorized
			}

			// Extract and assert JWT claims
			sub, okSub := claims["sub"].(string) // ✅ ensure it's a string (UUID)
			roleFloat, okRole := claims["role"].(float64)

			if !okSub || !okRole {
				log.Println("❌ JWT missing or invalid 'sub' or 'role'")
				return fiber.ErrUnauthorized
			}

			role := int(roleFloat)
			log.Printf("✅ JWT role: %d, sub: %s", role, sub)

			// Store in context
			ctx.Locals("user_id", sub)
			ctx.Locals("user", sub)
			ctx.Locals("role", role)

			// Special access mapping
			pegawai := []int{1, 1337, 2, 3, 4001, 4002, 4003, 4004, 5001}

			// Allow all roles if 0
			if len(roles) == 1 && roles[0] == 0 {
				return ctx.Next()
			}

			for _, r := range roles {
				if role == r {
					return ctx.Next()
				}
				if r == 2 {
					for _, p := range pegawai {
						if role == p {
							return ctx.Next()
						}
					}
				}
			}

			log.Printf("❌ Forbidden: role %d not allowed", role)
			return fiber.ErrForbidden
		},

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Printf("❌ JWT error: %v", err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access: " + err.Error(),
			})
		},

		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS512.Alg(),
			Key:    []byte(secret),
		},

		ContextKey: "jwt",
	})
}
