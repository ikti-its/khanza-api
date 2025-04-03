// internal/modules/rujukan/internal/repository/rujukan_keluar_repository.go

package repository

import "github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/entity"

type RujukanKeluarRepository interface {
	Insert(rujukan *entity.RujukanKeluar) error
	FindAll() ([]entity.RujukanKeluar, error)
	FindByNomorRawat(nomorRawat string) (entity.RujukanKeluar, error)
	FindByNomorRM(nomorRM string) ([]entity.RujukanKeluar, error)
	FindByTanggalRujuk(tanggal string) ([]entity.RujukanKeluar, error)
	Update(rujukan *entity.RujukanKeluar) error
	Delete(nomorRawat string) error
}
