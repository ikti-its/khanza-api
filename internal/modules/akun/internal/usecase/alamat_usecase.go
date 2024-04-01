package usecase

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/helper"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/repository"
)

type AlamatUseCase struct {
	Repository repository.AlamatRepository
}

func NewAlamatUseCase(repository *repository.AlamatRepository) *AlamatUseCase {
	return &AlamatUseCase{
		Repository: *repository,
	}
}

func (u *AlamatUseCase) Create(request *model.AlamatRequest, user string) model.AlamatResponse {
	updater := helper.MustParse(user)
	alamat := entity.Alamat{
		IdAkun:    helper.MustParse(request.IdAkun),
		Alamat:    request.Alamat,
		AlamatLat: request.AlamatLat,
		AlamatLon: request.AlamatLon,
		Kota:      request.Kota,
		KodePos:   request.KodePos,
		Updater:   updater,
	}

	if err := u.Repository.Insert(&alamat); err != nil {
		exception.PanicIfError(err, "Failed to create alamat")
	}

	response := model.AlamatResponse{
		IdAkun:    alamat.IdAkun.String(),
		Alamat:    alamat.Alamat,
		AlamatLat: alamat.AlamatLat,
		AlamatLon: alamat.AlamatLon,
		Kota:      alamat.Kota,
		KodePos:   alamat.KodePos,
	}

	return response
}

func (u *AlamatUseCase) GetById(id string) model.AlamatResponse {
	alamat, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alamat not found",
		})
	}

	response := model.AlamatResponse{
		IdAkun:    alamat.IdAkun.String(),
		Alamat:    alamat.Alamat,
		AlamatLat: alamat.AlamatLat,
		AlamatLon: alamat.AlamatLon,
		Kota:      alamat.Kota,
		KodePos:   alamat.KodePos,
	}

	return response
}

func (u *AlamatUseCase) Update(request *model.AlamatRequest, id, user string) model.AlamatResponse {
	alamat, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alamat not found",
		})
	}

	alamat.IdAkun = helper.MustParse(request.IdAkun)
	alamat.Alamat = request.Alamat
	alamat.AlamatLat = request.AlamatLat
	alamat.AlamatLon = request.AlamatLon
	alamat.Kota = request.Kota
	alamat.KodePos = request.KodePos
	alamat.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&alamat); err != nil {
		exception.PanicIfError(err, "Failed to update alamat")
	}

	response := model.AlamatResponse{
		IdAkun:    alamat.IdAkun.String(),
		Alamat:    alamat.Alamat,
		AlamatLat: alamat.AlamatLat,
		AlamatLon: alamat.AlamatLon,
		Kota:      alamat.Kota,
		KodePos:   alamat.KodePos,
	}

	return response
}

func (u *AlamatUseCase) Delete(id string) {
	alamat, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alamat not found",
		})
	}

	if err := u.Repository.Delete(&alamat); err != nil {
		exception.PanicIfError(err, "Failed to delete alamat")
	}
}
