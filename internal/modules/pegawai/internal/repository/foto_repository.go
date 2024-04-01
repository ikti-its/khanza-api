package repository

import (
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/entity"
	"github.com/google/uuid"
)

type FotoRepository interface {
	Insert(foto *entity.Foto) error
	FindAkunIdById(id uuid.UUID) (uuid.UUID, error)
	FindById(id uuid.UUID) (entity.Foto, error)
	Update(foto *entity.Foto) error
	Delete(foto *entity.Foto) error
}
