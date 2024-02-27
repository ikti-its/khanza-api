package repository

import (
	"github.com/fathoor/simkes-api/internal/akun/entity"
	"gorm.io/gorm"
	"math"
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

func (repository *akunRepositoryImpl) FindPage(page, size int) ([]entity.Akun, int, error) {
	var akun []entity.Akun
	var total int64

	if err := repository.DB.Model(&entity.Akun{}).Count(&total).Error; err != nil {
		return akun, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	err := repository.DB.Limit(size).Offset((page - 1) * size).Find(&akun).Error

	return akun, totalPage, err
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

func NewAkunRepositoryProvider(db *gorm.DB) AkunRepository {
	return &akunRepositoryImpl{db}
}
