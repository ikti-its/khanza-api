package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/entity"
)

type KamarRepository interface {
	Insert(kamar *entity.Kamar) error
	FindAll() ([]entity.Kamar, error)
	FindByNomorBed(nomorReg string) (entity.Kamar, error)
	FindByKodeKamar(nomorRM string) (entity.Kamar, error)
	Update(kamar *entity.Kamar) error
	Delete(nomorReg string) error
	GetAvailableRooms() ([]entity.Kamar, error)
}
