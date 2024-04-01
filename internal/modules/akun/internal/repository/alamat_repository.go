package repository

import (
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/entity"
	"github.com/google/uuid"
)

type AlamatRepository interface {
	Insert(alamat *entity.Alamat) error
	FindById(id uuid.UUID) (entity.Alamat, error)
	Update(alamat *entity.Alamat) error
	Delete(alamat *entity.Alamat) error
}
