package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/rawatinap/internal/entity"
)

type RawatInapRepository interface {
	Insert(rawatInap *entity.RawatInap) error
	FindAll() ([]entity.RawatInap, error)
	FindByNomorRawat(nomorRawat string) (entity.RawatInap, error)
	FindByTanggalMasuk(tanggal string) ([]entity.RawatInap, error)
	FindByNomorRM(nomorRM string) ([]entity.RawatInap, error)
	Update(rawatInap *entity.RawatInap) error
	Delete(nomorRawat string) error
}
