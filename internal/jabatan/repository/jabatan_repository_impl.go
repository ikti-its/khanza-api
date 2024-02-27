package repository

import (
	"github.com/fathoor/simkes-api/internal/jabatan/entity"
	"gorm.io/gorm"
)

type jabatanRepositoryImpl struct {
	*gorm.DB
}

func (repository *jabatanRepositoryImpl) Insert(jabatan *entity.Jabatan) error {
	return repository.DB.Create(&jabatan).Error
}

func (repository *jabatanRepositoryImpl) FindAll() ([]entity.Jabatan, error) {
	var jabatan []entity.Jabatan
	err := repository.DB.Find(&jabatan).Error

	return jabatan, err
}

func (repository *jabatanRepositoryImpl) FindByJabatan(j string) (entity.Jabatan, error) {
	var jabatan entity.Jabatan
	err := repository.DB.Take(&jabatan, "nama = ?", j).Error

	return jabatan, err
}

func (repository *jabatanRepositoryImpl) Update(jabatan *entity.Jabatan) error {
	return repository.DB.Save(&jabatan).Error
}

func (repository *jabatanRepositoryImpl) Delete(jabatan *entity.Jabatan) error {
	return repository.DB.Delete(&jabatan).Error
}

func NewJabatanRepositoryProvider(db *gorm.DB) JabatanRepository {
	return &jabatanRepositoryImpl{db}
}
