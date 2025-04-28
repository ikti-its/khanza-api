package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/entity"
)

type ResepPulangRepository interface {
	Insert(data *entity.ResepPulang) error
	FindAll() ([]entity.ResepPulang, error)
	FindByNoRawat(noRawat string) ([]entity.ResepPulang, error)
	FindByCompositeKey(noRawat, kodeBrng string, tanggal string, jam string) (*entity.ResepPulang, error)
	Update(data *entity.ResepPulang) error
	Delete(noRawat, kodeBrng, tanggal, jam string) error
}
