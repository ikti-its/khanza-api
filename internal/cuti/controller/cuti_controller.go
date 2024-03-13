package controller

import "github.com/gofiber/fiber/v2"

type CutiController interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
