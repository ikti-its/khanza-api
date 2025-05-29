package repository

import "github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"

type CatatanObservasiRanapKebidananRepository interface {
	Insert(observasi *entity.CatatanObservasiRanapKebidanan) error
	FindAll() ([]entity.CatatanObservasiRanapKebidanan, error)
	FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanapKebidanan, error)
	FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanapKebidanan, error)
	FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanapKebidanan, error)
	Update(observasi *entity.CatatanObservasiRanapKebidanan) error
	Delete(noRawat string, tglPerawatan string, jamRawat string) error
}
