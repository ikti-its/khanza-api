package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type SatuanUseCase struct {
	Repository repository.SatuanRepository
}

func NewSatuanUseCase(repository *repository.SatuanRepository) *SatuanUseCase {
	return &SatuanUseCase{
		Repository: *repository,
	}
}

func (u *SatuanUseCase) Get() []model.SatuanResponse {
	satuan, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all satuan")

	response := make([]model.SatuanResponse, len(satuan))
	for i, satuan := range satuan {
		response[i] = model.SatuanResponse{
			Id:   satuan.Id,
			Nama: satuan.Nama,
		}
	}

	return response
}
