package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/entity"
)

type PermintaanResepPulangRepository interface {
	Insert(data *entity.PermintaanResepPulang) error
	FindAll() ([]entity.PermintaanResepPulang, error)
	FindByNoRawat(noRawat string) ([]entity.PermintaanResepPulang, error)
	FindByNoPermintaan(noPermintaan string) (*entity.PermintaanResepPulang, error)
	Update(data *entity.PermintaanResepPulang) error
	Delete(noPermintaan string) error
}
