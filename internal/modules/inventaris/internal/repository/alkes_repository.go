package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type AlkesRepository interface {
	Insert(alkes *entity.Alkes) error
	Find() ([]entity.Alkes, error)
	FindPage(page, size int) ([]entity.Alkes, int, error)
	FindById(id uuid.UUID) (entity.Alkes, error)
	Update(alkes *entity.Alkes) error
	Delete(alkes *entity.Alkes) error
}
