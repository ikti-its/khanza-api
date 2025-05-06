package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/entity"
)

// interface in internal/modules/reseppulang/internal/repository/repository.go (or similar)
type PermintaanResepPulangRepository interface {
	InsertMany([]*entity.PermintaanResepPulang) error
	FindAll() ([]entity.PermintaanResepPulang, error)
	FindByNoRawat(noRawat string) ([]entity.PermintaanResepPulang, error)
	FindByNoPermintaan(noPermintaan string) (*entity.PermintaanResepPulang, error)
	Update(*entity.PermintaanResepPulang) error
	Delete(noPermintaan string) error
	GetByNoPermintaan(noPermintaan string) ([]entity.PermintaanResepPulang, error)
	GetByNoPermintaanWithHarga(noPermintaan string) ([]entity.ResepPulangObat, error)
}
