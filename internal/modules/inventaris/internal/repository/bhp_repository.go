package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type BhpRepository interface {
	Insert(bhp *entity.Bhp) error
	Find() ([]entity.Bhp, error)
	FindPage(page, size int) ([]entity.Bhp, int, error)
	FindById(id uuid.UUID) (entity.Bhp, error)
	Update(bhp *entity.Bhp) error
	Delete(bhp *entity.Bhp) error
}
