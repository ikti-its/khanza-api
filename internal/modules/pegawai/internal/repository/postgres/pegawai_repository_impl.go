package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/pegawai/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/pegawai/internal/repository"
	"gorm.io/gorm"
	"math"
)

type pegawaiRepositoryImpl struct {
	DB *gorm.DB
}

func NewPegawaiRepository(db *gorm.DB) repository.PegawaiRepository {
	return &pegawaiRepositoryImpl{db}
}

func (r *pegawaiRepositoryImpl) Insert(pegawai *entity.Pegawai) error {
	return r.DB.Table("pegawai").Create(&pegawai).Error
}

func (r *pegawaiRepositoryImpl) Find() ([]entity.Pegawai, error) {
	var pegawai []entity.Pegawai

	err := r.DB.Table("pegawai").
		Select("id, id_akun, nip, nama, jenis_kelamin, id_jabatan, id_departemen, id_status_aktif, jenis_pegawai, telepon, tanggal_masuk").
		Find(&pegawai).Error

	return pegawai, err
}

func (r *pegawaiRepositoryImpl) FindPage(page, size int) ([]entity.Pegawai, int, error) {
	var pegawai []entity.Pegawai
	var total int64

	if err := r.DB.Table("pegawai").Count(&total).Error; err != nil {
		return pegawai, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	err := r.DB.Table("pegawai").
		Select("id, id_akun, nip, nama, jenis_kelamin, id_jabatan, id_departemen, id_status_aktif, jenis_pegawai, telepon, tanggal_masuk").
		Limit(size).Offset(offset).Find(&pegawai).Error

	return pegawai, totalPage, err
}

func (r *pegawaiRepositoryImpl) FindById(id uuid.UUID) (entity.Pegawai, error) {
	var pegawai entity.Pegawai

	err := r.DB.Table("pegawai").
		Select("id, id_akun, nip, nama, jenis_kelamin, id_jabatan, id_departemen, id_status_aktif, jenis_pegawai, telepon, tanggal_masuk").
		Where("id = ?", id).
		First(&pegawai).Error

	return pegawai, err
}

func (r *pegawaiRepositoryImpl) Update(pegawai *entity.Pegawai) error {
	return r.DB.Table("pegawai").Save(&pegawai).Error
}

func (r *pegawaiRepositoryImpl) Delete(pegawai *entity.Pegawai) error {
	return r.DB.Table("pegawai").Delete(&pegawai).Error
}
