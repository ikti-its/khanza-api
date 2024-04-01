package repository

import (
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/entity"
	"github.com/google/uuid"
)

type JadwalRepository interface {
	Find() ([]entity.Jadwal, error)
	FindByHariId(id int) ([]entity.Jadwal, error)
	FindByPegawaiId(id uuid.UUID) ([]entity.Jadwal, error)
	FindById(id uuid.UUID) (entity.Jadwal, error)
	Update(jadwal *entity.Jadwal) error
}
