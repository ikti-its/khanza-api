package repository

import (
	"github.com/fathoor/simkes-api/internal/kehadiran/entity"
	"github.com/google/uuid"
)

type KehadiranRepository interface {
	Insert(kehadiran *entity.Kehadiran) error
	FindAll() ([]entity.Kehadiran, error)
	FindByNIP(nip string) ([]entity.Kehadiran, error)
	FindByID(id uuid.UUID) (entity.Kehadiran, error)
	FindLatestByNIP(nip string) (entity.Kehadiran, error)
	Update(kehadiran *entity.Kehadiran) error
	Delete(kehadiran *entity.Kehadiran) error
}
