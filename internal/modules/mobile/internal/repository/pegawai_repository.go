package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/entity"
)

type PegawaiRepository interface {
	Find() ([]entity.Pegawai, error)
	FindPage(page, size int) ([]entity.Pegawai, int, error)
	FindById(id uuid.UUID) (entity.Pegawai, error)
}
