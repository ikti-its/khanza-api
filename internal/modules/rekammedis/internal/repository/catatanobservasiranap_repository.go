package repository

import "github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"

type CatatanObservasiRanapRepository interface {
	Insert(observasi *entity.CatatanObservasiRanap) error
	FindAll() ([]entity.CatatanObservasiRanap, error)
	FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanap, error)
	FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanap, error)
	FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanap, error)
	Update(observasi *entity.CatatanObservasiRanap) error
	Delete(noRawat string, tglPerawatan string, jamRawat string) error
}
