package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type jadwalRepositoryImpl struct {
	DB *gorm.DB
}

func NewJadwalRepository(db *gorm.DB) repository.JadwalRepository {
	return &jadwalRepositoryImpl{DB: db}
}

func (r *jadwalRepositoryImpl) Find() ([]entity.Jadwal, error) {
	var jadwal []entity.Jadwal

	err := r.DB.Table("jadwal_pegawai jp").
		Select("jp.id, jp.id_pegawai, jp.id_hari, jp.id_shift, s.jam_masuk, s.jam_pulang").
		Joins("JOIN ref.shift s ON jp.id_shift = s.id").
		Order("jp.id_hari").
		Find(&jadwal).Error

	return jadwal, err
}

func (r *jadwalRepositoryImpl) FindByHariId(id int) ([]entity.Jadwal, error) {
	var jadwal []entity.Jadwal

	err := r.DB.Table("jadwal_pegawai jp").
		Select("jp.id, jp.id_pegawai, jp.id_hari, jp.id_shift, s.jam_masuk, s.jam_pulang").
		Joins("JOIN ref.shift s ON jp.id_shift = s.id").
		Where("jp.id_hari = ?", id).
		Order("jp.id_pegawai").
		Find(&jadwal).Error

	return jadwal, err
}

func (r *jadwalRepositoryImpl) FindByPegawaiId(id uuid.UUID) ([]entity.Jadwal, error) {
	var jadwal []entity.Jadwal

	err := r.DB.Table("jadwal_pegawai jp").
		Select("jp.id, jp.id_pegawai, jp.id_hari, jp.id_shift, s.jam_masuk, s.jam_pulang").
		Joins("JOIN ref.shift s ON jp.id_shift = s.id").
		Where("jp.id_pegawai = ?", id).
		Order("jp.id_hari").
		Find(&jadwal).Error

	return jadwal, err
}

func (r *jadwalRepositoryImpl) FindById(id uuid.UUID) (entity.Jadwal, error) {
	var jadwal entity.Jadwal

	err := r.DB.Table("jadwal_pegawai jp").
		Select("jp.id, jp.id_pegawai, jp.id_hari, jp.id_shift, s.jam_masuk, s.jam_pulang").
		Joins("JOIN ref.shift s ON jp.id_shift = s.id").
		Where("jp.id = ?", id).
		Order("jp.id_hari").
		First(&jadwal).Error

	return jadwal, err
}

func (r *jadwalRepositoryImpl) Update(jadwal *entity.Jadwal) error {
	return r.DB.Table("jadwal_pegawai").Omit("jam_masuk", "jam_pulang").Save(&jadwal).Error
}
