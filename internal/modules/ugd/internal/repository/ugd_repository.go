package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/ugd/internal/entity"
)

type UGDRepository interface {
	Insert(ugd *entity.UGD) error
	FindAll() ([]entity.UGD, error)
	FindByNomorReg(nomorReg string) (entity.UGD, error)
	FindByTanggal(tanggal string) ([]entity.UGD, error)
	FindByNomorRM(nomorRM string) (entity.UGD, error)
	Update(ugd *entity.UGD) error
	Delete(nomorReg string) error

	CheckDokterExists(kodeDokter string) (bool, error)
}
