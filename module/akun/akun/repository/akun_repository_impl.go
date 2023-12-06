package repository

import (
	"github.com/fathoor/simkes-api/module/akun/akun/entity"
	"gorm.io/gorm"
)

type akunRepositoryImpl struct {
	*gorm.DB
}

func (repository *akunRepositoryImpl) Insert(akun *entity.Akun) error {
	return repository.DB.Create(&akun).Error
}

func (repository *akunRepositoryImpl) FindAll() ([]entity.Akun, error) {
	var akun []entity.Akun
	err := repository.DB.Find(&akun).Error

	return akun, err
}

func (repository *akunRepositoryImpl) FindByNIP(nip string) (entity.Akun, error) {
	var akun entity.Akun
	err := repository.DB.Take(&akun, "nip = ?", nip).Error

	return akun, err
}

func (repository *akunRepositoryImpl) Update(akun *entity.Akun) error {
	return repository.DB.Save(&akun).Error
}

func (repository *akunRepositoryImpl) Delete(akun *entity.Akun) error {
	return repository.DB.Delete(&akun).Error
}

func ProvideAkunRepository(db *gorm.DB) AkunRepository {
	return &akunRepositoryImpl{db}
}
