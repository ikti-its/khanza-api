package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/entity"
)

type KehadiranRepository interface {
	Insert(kehadiran *entity.Kehadiran) error
	Find() ([]entity.Kehadiran, error)
	FindByPegawaiId(id uuid.UUID) ([]entity.Kehadiran, error)
	FindByTanggal(tanggal string) ([]entity.Kehadiran, error)
	FindById(id uuid.UUID) (entity.Kehadiran, error)
	Update(kehadiran *entity.Kehadiran) error
}
