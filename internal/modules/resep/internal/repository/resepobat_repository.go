package repository

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
)

type ResepObatRepository interface {
	Insert(c *fiber.Ctx, p *entity.ResepObat) error
	FindAll() ([]entity.ResepObat, error)
	FindByNoResep(noResep string) (*entity.ResepObat, error)
	Update(c *fiber.Ctx, p *entity.ResepObat) error
	Delete(c *fiber.Ctx, noResep string) error
	GetByNomorRawat(nomorRawat string) ([]entity.ResepObat, error)
	UpdateValidasi(ctx context.Context, noResep string, validasi bool) error
}
