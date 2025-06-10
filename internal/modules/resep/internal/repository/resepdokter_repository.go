package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
)

type ResepDokterRepository interface {
	Insert(resep *entity.ResepDokter) error
	FindAll() ([]entity.ResepDokter, error)
	FindByNoResep(noResep string) ([]entity.ResepDokter, error)
	Update(resep *entity.ResepDokter) error
	Delete(noResep string, kodeBarang string) error
}
