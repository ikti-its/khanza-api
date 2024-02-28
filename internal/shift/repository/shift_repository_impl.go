package repository

import (
	"github.com/fathoor/simkes-api/internal/shift/entity"
	"gorm.io/gorm"
)

type shiftRepositoryImpl struct {
	*gorm.DB
}

func (service *shiftRepositoryImpl) Insert(shift *entity.Shift) error {
	return service.DB.Create(&shift).Error
}

func (service *shiftRepositoryImpl) FindAll() ([]entity.Shift, error) {
	var shift []entity.Shift
	err := service.DB.Find(&shift).Error

	return shift, err
}

func (service *shiftRepositoryImpl) FindByNama(n string) (entity.Shift, error) {
	var shift entity.Shift
	err := service.DB.Take(&shift, "nama = ?", n).Error

	return shift, err
}

func (service *shiftRepositoryImpl) Update(shift *entity.Shift) error {
	return service.DB.Save(&shift).Error
}

func (service *shiftRepositoryImpl) Delete(shift *entity.Shift) error {
	return service.DB.Delete(&shift).Error
}

func NewShiftRepositoryProvider(db *gorm.DB) ShiftRepository {
	return &shiftRepositoryImpl{db}
}
