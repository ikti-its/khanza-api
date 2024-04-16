package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type MedisRepository interface {
	Insert(medis *entity.Medis) error
	Find() ([]entity.Medis, error)
	FindPage(page, size int) ([]entity.Medis, int, error)
	FindByJenis(jenis string) ([]entity.Medis, error)
	FindById(id uuid.UUID) (entity.Medis, error)
	Update(medis *entity.Medis) error
	Delete(medis *entity.Medis) error
}
