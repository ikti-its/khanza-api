package repository

import (
	"context"

	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
)

type ResepObatRepository interface {
	Insert(resep *entity.ResepObat) error
	FindAll() ([]entity.ResepObat, error)
	FindByNoResep(noResep string) (*entity.ResepObat, error)
	Update(resep *entity.ResepObat) error
	Delete(noResep string) error
	GetByNomorRawat(nomorRawat string) ([]entity.ResepObat, error)
	UpdateValidasi(ctx context.Context, noResep string, validasi bool) error
}
