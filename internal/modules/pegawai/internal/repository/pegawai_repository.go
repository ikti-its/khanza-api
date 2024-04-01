package repository

import (
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/entity"
	"github.com/google/uuid"
)

type PegawaiRepository interface {
	Insert(pegawai *entity.Pegawai) error
	Find() ([]entity.Pegawai, error)
	FindPage(page, size int) ([]entity.Pegawai, int, error)
	FindById(id uuid.UUID) (entity.Pegawai, error)
	Update(pegawai *entity.Pegawai) error
	Delete(pegawai *entity.Pegawai) error
}
