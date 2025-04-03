package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/entity"
)

type RujukanMasukRepository interface {
	Insert(rujukan *entity.RujukanMasuk) error
	FindAll() ([]entity.RujukanMasuk, error)
	FindByNomorRawat(nomorRawat string) (entity.RujukanMasuk, error)
	FindByNomorRM(nomorRM string) ([]entity.RujukanMasuk, error)
	FindByTanggalMasuk(tanggal string) ([]entity.RujukanMasuk, error)
	Update(rujukan *entity.RujukanMasuk) error
	Delete(nomorRawat string) error
}
