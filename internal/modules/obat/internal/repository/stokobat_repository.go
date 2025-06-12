package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/entity"
)

type GudangBarangRepository interface {
	Insert(barang *entity.GudangBarang) error
	FindAll() ([]entity.GudangBarang, error)
	FindByID(id string) (*entity.GudangBarang, error)
	Update(barang *entity.GudangBarang) error
	Delete(id string) error
	FindByIDBarangMedis(idBarangMedis string) ([]entity.GudangBarang, error)
}
