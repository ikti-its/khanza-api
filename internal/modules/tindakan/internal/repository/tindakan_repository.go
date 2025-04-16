package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/entity"
)

type TindakanRepository interface {
	Insert(tindakan *entity.Tindakan) error
	FindAll() ([]entity.Tindakan, error)
	FindByNomorRawat(nomorRawat string) ([]entity.Tindakan, error)
	Update(tindakan *entity.Tindakan) error
	Delete(nomorRawat string, jamRawat string) error
	GetAllJenisTindakan() ([]entity.JenisTindakan, error)

	CheckDokterExists(kodeDokter string) (bool, error)
}
