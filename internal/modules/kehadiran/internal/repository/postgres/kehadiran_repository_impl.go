package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type kehadiranRepositoryImpl struct {
	DB *gorm.DB
}

func NewKehadiranRepository(db *gorm.DB) repository.KehadiranRepository {
	return &kehadiranRepositoryImpl{DB: db}
}

func (r *kehadiranRepositoryImpl) Insert(kehadiran *entity.Kehadiran) error {
	return r.DB.Table("kehadiran").Create(&kehadiran).Error
}

func (r *kehadiranRepositoryImpl) Find() ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran

	err := r.DB.Table("kehadiran").
		Select("id, id_pegawai, tanggal, jam_masuk, jam_pulang, keterangan").
		Order("tanggal").
		Find(&kehadiran).Error

	return kehadiran, err
}

func (r *kehadiranRepositoryImpl) FindByPegawaiId(id uuid.UUID) ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran

	err := r.DB.Table("kehadiran").
		Select("id, id_pegawai, tanggal, jam_masuk, jam_pulang, keterangan").
		Where("id_pegawai = ?", id).
		Order("tanggal").
		Find(&kehadiran).Error

	return kehadiran, err
}

func (r *kehadiranRepositoryImpl) FindByTanggal(tanggal string) ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran

	err := r.DB.Table("kehadiran").
		Select("id, id_pegawai, tanggal, jam_masuk, jam_pulang, keterangan").
		Where("tanggal = ?", tanggal).
		Order("tanggal").
		Find(&kehadiran).Error

	return kehadiran, err
}

func (r *kehadiranRepositoryImpl) FindById(id uuid.UUID) (entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran

	err := r.DB.Table("kehadiran").
		Select("id, id_pegawai, tanggal, jam_masuk, jam_pulang, keterangan").
		Where("id = ?", id).
		First(&kehadiran).Error

	return kehadiran, err
}

func (r *kehadiranRepositoryImpl) Update(kehadiran *entity.Kehadiran) error {
	return r.DB.Table("kehadiran").Save(&kehadiran).Error
}
