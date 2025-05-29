package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
)

type ResepDokterRacikanRepository interface {
	Insert(racikan *entity.ResepDokterRacikan) error
	FindAll() ([]entity.ResepDokterRacikan, error)
	FindByNoResep(noResep string) ([]entity.ResepDokterRacikan, error)
	Update(racikan *entity.ResepDokterRacikan) error
	Delete(noResep string, noRacik string) error
}
