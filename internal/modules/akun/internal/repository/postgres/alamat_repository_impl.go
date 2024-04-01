package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type alamatRepositoryImpl struct {
	DB *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) repository.AlamatRepository {
	return &alamatRepositoryImpl{db}
}

func (r *alamatRepositoryImpl) Insert(alamat *entity.Alamat) error {
	return r.DB.Table("alamat").Create(&alamat).Error
}

func (r *alamatRepositoryImpl) FindById(id uuid.UUID) (entity.Alamat, error) {
	var alamat entity.Alamat

	err := r.DB.Table("alamat").
		Select("id_akun, alamat, alamat_lat, alamat_lon, kota, kode_pos").
		Where("id_akun = ?", id).
		First(&alamat).Error

	return alamat, err
}

func (r *alamatRepositoryImpl) Update(alamat *entity.Alamat) error {
	return r.DB.Table("alamat").Where("id_akun = ?", alamat.IdAkun).Save(&alamat).Error
}

func (r *alamatRepositoryImpl) Delete(alamat *entity.Alamat) error {
	return r.DB.Table("alamat").Where("id_akun = ?", alamat.IdAkun).Delete(&alamat).Error
}
