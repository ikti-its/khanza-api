package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/model"
)

type TindakanRepository interface {
	Insert(tindakan *entity.Tindakan) error
	FindAll() ([]entity.Tindakan, error)
	FindByNomorRawat(nomorRawat string) ([]entity.Tindakan, error)
	Update(tindakan *entity.Tindakan) error
	Delete(nomorRawat string, jamRawat string) error
	GetAllJenisTindakan() ([]entity.JenisTindakan, error)
	FindJenisByKode(kode string) (*model.JenisTindakan, error)
	FindByNomorRawatAndJamRawat(nomorRawat, jamRawat string) (*entity.Tindakan, error)

	CheckDokterExists(kodeDokter string) (bool, error)
}
