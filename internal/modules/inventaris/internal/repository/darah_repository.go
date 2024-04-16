package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type DarahRepository interface {
	Insert(darah *entity.Darah) error
	Find() ([]entity.Darah, error)
	FindPage(page, size int) ([]entity.Darah, int, error)
	FindById(id uuid.UUID) (entity.Darah, error)
	Update(darah *entity.Darah) error
	Delete(darah *entity.Darah) error
}
