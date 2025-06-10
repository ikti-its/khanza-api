package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/entity"
)

type PemberianObatRepository interface {
	Insert(obat *entity.PemberianObat) error
	FindAll() ([]entity.PemberianObat, error)
	FindByNomorRawat(nomorRawat string) ([]entity.PemberianObat, error)
	Update(obat *entity.PemberianObat) error
	Delete(nomorRawat string, jamBeri string) error
	GetAllDataBarang() ([]entity.DataBarang, error)
}
