package controller

import "github.com/gofiber/fiber/v2"

type JadwalPegawaiController interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetByPK(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
