package repository

import (
	"github.com/fathoor/simkes-api/internal/jadwal-pegawai/entity"
	"gorm.io/gorm"
)

type jadwalPegawaiRepositoryImpl struct {
	*gorm.DB
}

func (repository *jadwalPegawaiRepositoryImpl) Insert(jadwalPegawai *entity.JadwalPegawai) error {
	return repository.DB.Create(&jadwalPegawai).Error
}

func (repository *jadwalPegawaiRepositoryImpl) FindAll() ([]entity.JadwalPegawai, error) {
	var jadwalPegawai []entity.JadwalPegawai
	err := repository.DB.Find(&jadwalPegawai).Error

	return jadwalPegawai, err
}

func (repository *jadwalPegawaiRepositoryImpl) FindByNIP(nip string) ([]entity.JadwalPegawai, error) {
	var jadwalPegawai []entity.JadwalPegawai
	err := repository.DB.Find(&jadwalPegawai, "nip = ?", nip).Error

	return jadwalPegawai, err
}

func (repository *jadwalPegawaiRepositoryImpl) FindByTahunBulan(tahun, bulan int16) ([]entity.JadwalPegawai, error) {
	var jadwalPegawai []entity.JadwalPegawai
	err := repository.DB.Find(&jadwalPegawai, "tahun = ? AND bulan = ?", tahun, bulan).Error

	return jadwalPegawai, err
}

func (repository *jadwalPegawaiRepositoryImpl) FindByPK(nip string, tahun, bulan, hari int16) (entity.JadwalPegawai, error) {
	var jadwalPegawai entity.JadwalPegawai
	err := repository.DB.Take(&jadwalPegawai, "nip = ? AND tahun = ? AND bulan = ? AND hari = ?", nip, tahun, bulan, hari).Error

	return jadwalPegawai, err
}

func (repository *jadwalPegawaiRepositoryImpl) Update(jadwalPegawai *entity.JadwalPegawai) error {
	return repository.DB.Save(&jadwalPegawai).Error
}

func (repository *jadwalPegawaiRepositoryImpl) Delete(jadwalPegawai *entity.JadwalPegawai) error {
	return repository.DB.Delete(&jadwalPegawai).Error
}

func NewJadwalPegawaiRepositoryProvider(db *gorm.DB) JadwalPegawaiRepository {
	return &jadwalPegawaiRepositoryImpl{db}
}
