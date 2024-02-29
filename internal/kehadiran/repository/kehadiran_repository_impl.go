package repository

import (
	"github.com/fathoor/simkes-api/internal/kehadiran/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type kehadiranRepositoryImpl struct {
	*gorm.DB
}

func (repository *kehadiranRepositoryImpl) Insert(kehadiran *entity.Kehadiran) error {
	return repository.DB.Create(&kehadiran).Error
}

func (repository *kehadiranRepositoryImpl) FindAll() ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran
	err := repository.DB.Preload("Shift").Find(&kehadiran).Error

	return kehadiran, err
}

func (repository *kehadiranRepositoryImpl) FindByNIP(nip string) ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran
	err := repository.DB.Preload("Shift").Find(&kehadiran, "nip = ?", nip).Error

	return kehadiran, err
}

func (repository *kehadiranRepositoryImpl) FindByID(id uuid.UUID) (entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran
	err := repository.DB.Preload("Shift").Take(&kehadiran, "id = ?", id).Error

	return kehadiran, err
}

func (repository *kehadiranRepositoryImpl) FindLatestByNIP(nip string) (entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran
	err := repository.DB.Preload("Shift").Order("tanggal desc").Take(&kehadiran, "nip = ?", nip).Error

	return kehadiran, err
}

func (repository *kehadiranRepositoryImpl) Update(kehadiran *entity.Kehadiran) error {
	return repository.DB.Save(&kehadiran).Error
}

func (repository *kehadiranRepositoryImpl) Delete(kehadiran *entity.Kehadiran) error {
	return repository.DB.Delete(&kehadiran).Error
}

func NewKehadiranRepositoryProvider(db *gorm.DB) KehadiranRepository {
	return &kehadiranRepositoryImpl{db}
}
