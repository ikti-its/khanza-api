package repository

import "github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"

type CatatanObservasiRanapPostpartumRepository interface {
	Insert(data *entity.CatatanObservasiRanapPostpartum) error
	FindAll() ([]entity.CatatanObservasiRanapPostpartum, error)
	FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanapPostpartum, error)
	FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanapPostpartum, error)
	FindByNoRawatAndTanggal(noRawat, tanggal string) ([]entity.CatatanObservasiRanapPostpartum, error)
	Update(data *entity.CatatanObservasiRanapPostpartum) error
	Delete(noRawat, tglPerawatan, jamRawat string) error
}
