package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
)

type akunRepositoryImpl struct {
	DB *gorm.DB
}

func NewAkunRepository(db *gorm.DB) repository.AkunRepository {
	return &akunRepositoryImpl{db}
}

func (r *akunRepositoryImpl) Insert(akun *entity.Akun) error {
	return r.DB.Table("akun").Create(&akun).Error
}

func (r *akunRepositoryImpl) Find() ([]entity.Akun, error) {
	var akun []entity.Akun

	err := r.DB.Table("akun").Select("id, email, foto, role").Find(&akun).Error

	return akun, err
}

func (r *akunRepositoryImpl) FindPage(page, size int) ([]entity.Akun, int, error) {
	var akun []entity.Akun
	var total int64

	if err := r.DB.Table("akun").Count(&total).Error; err != nil {
		return akun, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	err := r.DB.Table("akun").Select("id, email, foto, role").Find(&akun).Limit(size).Offset(offset).Error

	return akun, totalPage, err
}

func (r *akunRepositoryImpl) FindById(id uuid.UUID) (entity.Akun, error) {
	var akun entity.Akun

	err := r.DB.Table("akun").Select("id, email, foto, role").Where("id = ?", id).First(&akun).Error

	return akun, err
}

func (r *akunRepositoryImpl) Update(akun *entity.Akun) error {
	return r.DB.Table("akun").Save(&akun).Error
}

func (r *akunRepositoryImpl) Delete(akun *entity.Akun) error {
	return r.DB.Table("akun").Delete(&akun).Error
}
