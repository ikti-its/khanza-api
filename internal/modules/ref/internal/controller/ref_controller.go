package controller

import (
	"github.com/gofiber/fiber/v2"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/ref/internal/usecase"
)

type RefController struct {
	UseCase *usecase.RefUseCase
}

func NewRefController(useCase *usecase.RefUseCase) *RefController {
	return &RefController{
		UseCase: useCase,
	}
}

func (c *RefController) GetRole(ctx *fiber.Ctx) error {
	response := c.UseCase.GetRole()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetJabatan(ctx *fiber.Ctx) error {
	response := c.UseCase.GetJabatan()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetDepartemen(ctx *fiber.Ctx) error {
	response := c.UseCase.GetDepartemen()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetStatusAktif(ctx *fiber.Ctx) error {
	response := c.UseCase.GetStatusAktif()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetShift(ctx *fiber.Ctx) error {
	response := c.UseCase.GetShift()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetAlasanCuti(ctx *fiber.Ctx) error {
	response := c.UseCase.GetAlasanCuti()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RefController) GetIndustriFarmasi(ctx *fiber.Ctx) error {
	response := c.UseCase.GetIndustriFarmasi()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
func (c *RefController) GetSatuanBarangMedis(ctx *fiber.Ctx) error {
	response := c.UseCase.GetSatuanBarangMedis()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
func (c *RefController) GetJenisBarangMedis(ctx *fiber.Ctx) error {
	response := c.UseCase.GetJenisBarangMedis()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
func (c *RefController) GetKategoriBarangMedis(ctx *fiber.Ctx) error {
	response := c.UseCase.GetKategoriBarangMedis()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
func (c *RefController) GetGolonganBarangMedis(ctx *fiber.Ctx) error {
	response := c.UseCase.GetGolonganBarangMedis()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
func (c *RefController) GetRuangan(ctx *fiber.Ctx) error {
	response := c.UseCase.GetRuangan()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
func (c *RefController) GetSupplierBarangMedis(ctx *fiber.Ctx) error {
	response := c.UseCase.GetSupplierBarangMedis()

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}
