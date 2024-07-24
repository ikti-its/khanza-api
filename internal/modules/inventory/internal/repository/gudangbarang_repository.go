package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/entity"
)

type GudangBarangRepository interface {
	Insert(opname *entity.GudangBarang) error
	Find() ([]entity.GudangBarang, error)
	FindById(id uuid.UUID) (entity.GudangBarang, error)
	Update(opname *entity.GudangBarang) error
	Delete(opname *entity.GudangBarang) error
}
