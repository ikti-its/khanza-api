package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/akun/internal/entity"
)

type AlamatRepository interface {
	Insert(alamat *entity.Alamat) error
	FindById(id uuid.UUID) (entity.Alamat, error)
	Update(alamat *entity.Alamat) error
	Delete(alamat *entity.Alamat) error
}
