package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Login(c *fiber.Ctx) error
}
