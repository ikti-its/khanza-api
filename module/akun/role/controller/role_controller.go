package controller

import "github.com/gofiber/fiber/v2"

type RoleController interface {
	Route(app *fiber.App)
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
