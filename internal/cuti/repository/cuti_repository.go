package repository

import (
	"github.com/fathoor/simkes-api/internal/cuti/entity"
	"github.com/google/uuid"
)

type CutiRepository interface {
	Insert(cuti *entity.Cuti) error
	FindAll() ([]entity.Cuti, error)
	FindByNIP(n string) ([]entity.Cuti, error)
	FindByID(id uuid.UUID) (entity.Cuti, error)
	Update(cuti *entity.Cuti) error
	Delete(cuti *entity.Cuti) error
}
