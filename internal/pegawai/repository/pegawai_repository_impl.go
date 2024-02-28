package repository

import (
	"github.com/fathoor/simkes-api/internal/pegawai/entity"
	"gorm.io/gorm"
	"math"
)

type pegawaiRepositoryImpl struct {
	*gorm.DB
}

func (repository *pegawaiRepositoryImpl) Insert(pegawai *entity.Pegawai) error {
	return repository.DB.Create(&pegawai).Error
}

func (repository *pegawaiRepositoryImpl) FindAll() ([]entity.Pegawai, error) {
	var pegawai []entity.Pegawai
	err := repository.DB.Find(&pegawai).Error

	return pegawai, err
}

func (repository *pegawaiRepositoryImpl) FindPage(page, size int) ([]entity.Pegawai, int, error) {
	var pegawai []entity.Pegawai
	var total int64

	if err := repository.DB.Model(&entity.Pegawai{}).Count(&total).Error; err != nil {
		return pegawai, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	err := repository.DB.Limit(size).Offset((page - 1) * size).Find(&pegawai).Error

	return pegawai, totalPage, err
}

func (repository *pegawaiRepositoryImpl) FindByNIP(n string) (entity.Pegawai, error) {
	var pegawai entity.Pegawai
	err := repository.DB.Take(&pegawai, "nip = ?", n).Error

	return pegawai, err
}

func (repository *pegawaiRepositoryImpl) Update(pegawai *entity.Pegawai) error {
	return repository.DB.Save(&pegawai).Error
}

func (repository *pegawaiRepositoryImpl) Delete(pegawai *entity.Pegawai) error {
	return repository.DB.Delete(&pegawai).Error
}

func NewPegawaiRepositoryProvider(db *gorm.DB) PegawaiRepository {
	return &pegawaiRepositoryImpl{db}
}
