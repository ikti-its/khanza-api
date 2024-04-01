package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/pegawai/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type fotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewFotoRepository(db *gorm.DB) repository.FotoRepository {
	return &fotoRepositoryImpl{db}
}

func (r *fotoRepositoryImpl) Insert(foto *entity.Foto) error {
	return r.DB.Table("foto_pegawai").Create(&foto).Error
}

func (r *fotoRepositoryImpl) FindAkunIdById(id uuid.UUID) (uuid.UUID, error) {
	var record struct {
		Id uuid.UUID `gorm:"column:id_akun"`
	}

	err := r.DB.Table("pegawai").
		Select("id_akun").
		Where("id = ?", id).
		First(&record).Error

	return record.Id, err
}

func (r *fotoRepositoryImpl) FindById(id uuid.UUID) (entity.Foto, error) {
	var foto entity.Foto

	err := r.DB.Table("foto_pegawai").
		Select("id_pegawai, foto").
		Where("id_pegawai = ?", id).
		First(&foto).Error

	return foto, err
}

func (r *fotoRepositoryImpl) Update(foto *entity.Foto) error {
	return r.DB.Table("foto_pegawai").Where("id_pegawai = ?", foto.IdPegawai).Save(&foto).Error
}

func (r *fotoRepositoryImpl) Delete(foto *entity.Foto) error {
	return r.DB.Table("foto_pegawai").Where("id_pegawai = ?", foto.IdPegawai).Delete(&foto).Error
}
