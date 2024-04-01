package repository

import (
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/entity"
	"github.com/google/uuid"
)

type BerkasRepository interface {
	Insert(berkas *entity.Berkas) error
	FindAkunIdById(id uuid.UUID) (uuid.UUID, error)
	FindById(id uuid.UUID) (entity.Berkas, error)
	Update(berkas *entity.Berkas) error
	Delete(berkas *entity.Berkas) error
}
