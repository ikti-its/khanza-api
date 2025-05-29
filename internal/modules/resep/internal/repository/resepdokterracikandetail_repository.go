package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
)

type ResepDokterRacikanDetailRepository interface {
	Insert(detail *entity.ResepDokterRacikanDetail) error
	FindAll() ([]entity.ResepDokterRacikanDetail, error)
	FindByNoResepAndNoRacik(noResep, noRacik string) ([]entity.ResepDokterRacikanDetail, error)
	Update(detail *entity.ResepDokterRacikanDetail) error
	Delete(noResep, noRacik, kodeBrng string) error
}
