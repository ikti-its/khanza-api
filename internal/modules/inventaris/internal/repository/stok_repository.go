package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type StokRepository interface {
	Insert(stok *entity.Stok) error
	Find() ([]entity.Stok, error)
	FindPage(page, size int) ([]entity.Stok, int, error)
	FindById(id uuid.UUID) (entity.Stok, error)
	Update(stok *entity.Stok) error
	Delete(stok *entity.Stok) error
}
