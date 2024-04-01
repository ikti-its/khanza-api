package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type berkasRepositoryImpl struct {
	DB *gorm.DB
}

func NewBerkasRepository(db *gorm.DB) repository.BerkasRepository {
	return &berkasRepositoryImpl{db}
}

func (r *berkasRepositoryImpl) Insert(berkas *entity.Berkas) error {
	return r.DB.Table("berkas_pegawai").Create(&berkas).Error
}

func (r *berkasRepositoryImpl) FindAkunIdById(id uuid.UUID) (uuid.UUID, error) {
	var record struct {
		Id uuid.UUID `gorm:"column:id_akun"`
	}

	err := r.DB.Table("pegawai").
		Select("id_akun").
		Where("id = ?", id).
		First(&record).Error

	return record.Id, err
}

func (r *berkasRepositoryImpl) FindById(id uuid.UUID) (entity.Berkas, error) {
	var berkas entity.Berkas

	err := r.DB.Table("berkas_pegawai").
		Select("id_pegawai, nik, tempat_lahir, tanggal_lahir, agama, pendidikan, ktp, kk, npwp, bpjs, ijazah, skck, str, serkom").
		Where("id_pegawai = ?", id).
		First(&berkas).Error

	return berkas, err
}

func (r *berkasRepositoryImpl) Update(berkas *entity.Berkas) error {
	return r.DB.Table("berkas_pegawai").Where("id_pegawai = ?", berkas.IdPegawai).Save(&berkas).Error
}

func (r *berkasRepositoryImpl) Delete(berkas *entity.Berkas) error {
	return r.DB.Table("berkas_pegawai").Where("id_pegawai = ?", berkas.IdPegawai).Delete(&berkas).Error
}
