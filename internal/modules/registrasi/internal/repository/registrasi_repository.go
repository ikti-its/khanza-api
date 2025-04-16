package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/model"
)

type RegistrasiRepository interface {
	Insert(registrasi *entity.Registrasi) error
	FindAll() ([]entity.Registrasi, error)
	FindByNomorReg(nomorReg string) (entity.Registrasi, error)
	FindByTanggal(tanggal string) (entity.Registrasi, error)
	FindByNomorRM(nomorRM string) (entity.Registrasi, error)
	Update(registrasi *entity.Registrasi) error
	Delete(nomorReg string) error
	UpdateStatusKamar(nomorReg string, status string) error
	FindPendingRoomRequests() ([]entity.Registrasi, error)
	AssignKamar(nomorReg string, kamarID string) error
	GetAllDokter() ([]model.DokterResponse, error)
	GetNamaDokter(kode string) (string, error)

	CheckDokterExists(kodeDokter string) (bool, error)
}
