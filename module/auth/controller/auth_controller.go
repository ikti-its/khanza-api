package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Route(app *fiber.App)
	Login(c *fiber.Ctx) error
}
