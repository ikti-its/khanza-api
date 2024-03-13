package repository

import (
	"github.com/fathoor/simkes-api/internal/cuti/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cutiRepositoryImpl struct {
	*gorm.DB
}

func (repository *cutiRepositoryImpl) Insert(cuti *entity.Cuti) error {
	return repository.DB.Create(&cuti).Error
}

func (repository *cutiRepositoryImpl) FindAll() ([]entity.Cuti, error) {
	var cuti []entity.Cuti
	err := repository.DB.Preload("Pegawai").Find(&cuti).Error

	return cuti, err
}

func (repository *cutiRepositoryImpl) FindByNIP(n string) ([]entity.Cuti, error) {
	var cuti []entity.Cuti
	err := repository.DB.Preload("Pegawai").Find(&cuti, "nip = ?", n).Error

	return cuti, err
}

func (repository *cutiRepositoryImpl) FindByID(id uuid.UUID) (entity.Cuti, error) {
	var cuti entity.Cuti
	err := repository.DB.Preload("Pegawai").Take(&cuti, "id = ?", id).Error

	return cuti, err
}

func (repository *cutiRepositoryImpl) Update(cuti *entity.Cuti) error {
	return repository.DB.Save(&cuti).Error
}

func (repository *cutiRepositoryImpl) Delete(cuti *entity.Cuti) error {
	return repository.DB.Delete(&cuti).Error
}

func NewCutiRepositoryProvider(db *gorm.DB) CutiRepository {
	return &cutiRepositoryImpl{db}
}
