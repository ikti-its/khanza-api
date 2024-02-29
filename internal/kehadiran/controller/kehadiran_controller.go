package controller

import "github.com/gofiber/fiber/v2"

type KehadiranController interface {
	CheckIn(c *fiber.Ctx) error
	CheckOut(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
