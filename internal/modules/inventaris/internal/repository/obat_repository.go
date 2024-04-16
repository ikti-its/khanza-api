package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type ObatRepository interface {
	Insert(obat *entity.Obat) error
	Find() ([]entity.Obat, error)
	FindPage(page, size int) ([]entity.Obat, int, error)
	FindById(id uuid.UUID) (entity.Obat, error)
	Update(obat *entity.Obat) error
	Delete(obat *entity.Obat) error
}
