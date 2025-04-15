package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/dokterjaga/internal/entity"
)

type DokterJagaRepository interface {
	Insert(dokter *entity.DokterJaga) error
	FindAll() ([]entity.DokterJaga, error)
	FindByKodeDokter(kodeDokter string) ([]entity.DokterJaga, error)
	Update(dokter *entity.DokterJaga) error
	Delete(kodeDokter string, hariKerja string) error

	// Optional extensions (depending on use case)
	FindByStatus(status string) ([]entity.DokterJaga, error)
	UpdateStatus(kodeDokter string, hariKerja string, newStatus string) error
}
