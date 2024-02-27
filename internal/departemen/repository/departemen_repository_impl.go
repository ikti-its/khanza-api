package repository

import (
	"github.com/fathoor/simkes-api/internal/departemen/entity"
	"gorm.io/gorm"
)

type departemenRepositoryImpl struct {
	*gorm.DB
}

func (repository *departemenRepositoryImpl) Insert(departemen *entity.Departemen) error {
	return repository.DB.Create(&departemen).Error
}

func (repository *departemenRepositoryImpl) FindAll() ([]entity.Departemen, error) {
	var departemen []entity.Departemen
	err := repository.DB.Find(&departemen).Error

	return departemen, err
}

func (repository *departemenRepositoryImpl) FindByDepartemen(d string) (entity.Departemen, error) {
	var departemen entity.Departemen
	err := repository.DB.Take(&departemen, "nama = ?", d).Error

	return departemen, err
}

func (repository *departemenRepositoryImpl) Update(departemen *entity.Departemen) error {
	return repository.DB.Save(&departemen).Error
}

func (repository *departemenRepositoryImpl) Delete(departemen *entity.Departemen) error {
	return repository.DB.Delete(&departemen).Error
}

func NewDepartemenRepositoryProvider(db *gorm.DB) DepartemenRepository {
	return &departemenRepositoryImpl{db}
}
