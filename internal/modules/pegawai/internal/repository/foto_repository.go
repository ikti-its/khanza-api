package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/pegawai/internal/entity"
)

type FotoRepository interface {
	Insert(foto *entity.Foto) error
	FindAkunIdById(id uuid.UUID) (uuid.UUID, error)
	FindById(id uuid.UUID) (entity.Foto, error)
	Update(foto *entity.Foto) error
	Delete(foto *entity.Foto) error
}
