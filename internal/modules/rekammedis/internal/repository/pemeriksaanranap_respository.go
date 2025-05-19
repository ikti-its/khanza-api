package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
)

type PemeriksaanRanapRepository interface {
	Insert(pemeriksaan *entity.PemeriksaanRanap) error
	FindAll() ([]entity.PemeriksaanRanap, error)
	FindByNomorRawat(nomorRawat string) (entity.PemeriksaanRanap, error)
	FindByTanggal(tanggal string) ([]entity.PemeriksaanRanap, error)
	Update(pemeriksaan *entity.PemeriksaanRanap) error
	Delete(nomorRawat string) error
}
