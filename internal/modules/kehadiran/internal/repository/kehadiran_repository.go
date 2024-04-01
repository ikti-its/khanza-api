package repository

import (
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/entity"
	"github.com/google/uuid"
)

type KehadiranRepository interface {
	Insert(kehadiran *entity.Kehadiran) error
	Find() ([]entity.Kehadiran, error)
	FindByPegawaiId(id uuid.UUID) ([]entity.Kehadiran, error)
	FindByTanggal(tanggal string) ([]entity.Kehadiran, error)
	FindById(id uuid.UUID) (entity.Kehadiran, error)
	Update(kehadiran *entity.Kehadiran) error
}
