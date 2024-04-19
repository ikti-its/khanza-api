package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type DarahUseCase struct {
	Repository repository.DarahRepository
}

func NewDarahUseCase(repository *repository.DarahRepository) *DarahUseCase {
	return &DarahUseCase{
		Repository: *repository,
	}
}

func (u *DarahUseCase) Create(request *model.DarahRequest, user string) model.DarahResponse {
	updater := helper.MustParse(user)
	darah := entity.Darah{
		Id:          helper.MustNew(),
		IdMedis:     helper.MustParse(request.IdMedis),
		Keterangan:  request.Keterangan,
		Kadaluwarsa: helper.ParseTime(request.Kadaluwarsa, "2006-01-02"),
		Updater:     updater,
	}

	if err := u.Repository.Insert(&darah); err != nil {
		exception.PanicIfError(err, "Failed to insert darah")
	}

	response := model.DarahResponse{
		Id:          darah.Id.String(),
		IdMedis:     darah.IdMedis.String(),
		Keterangan:  darah.Keterangan,
		Kadaluwarsa: helper.FormatTime(darah.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *DarahUseCase) Get() []model.DarahResponse {
	darah, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all darah")

	response := make([]model.DarahResponse, len(darah))
	for i, darah := range darah {
		response[i] = model.DarahResponse{
			Id:          darah.Id.String(),
			IdMedis:     darah.IdMedis.String(),
			Keterangan:  darah.Keterangan,
			Kadaluwarsa: helper.FormatTime(darah.Kadaluwarsa, "2006-01-02"),
		}
	}

	return response
}

func (u *DarahUseCase) GetPage(page, size int) model.DarahPageResponse {
	darah, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged darah")

	response := make([]model.DarahResponse, len(darah))
	for i, darah := range darah {
		response[i] = model.DarahResponse{
			Id:          darah.Id.String(),
			IdMedis:     darah.IdMedis.String(),
			Keterangan:  darah.Keterangan,
			Kadaluwarsa: helper.FormatTime(darah.Kadaluwarsa, "2006-01-02"),
		}
	}

	pagedResponse := model.DarahPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Darah: response,
	}

	return pagedResponse
}

func (u *DarahUseCase) GetById(id string) model.DarahResponse {
	darah, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Darah not found",
		})
	}

	response := model.DarahResponse{
		Id:          darah.Id.String(),
		IdMedis:     darah.IdMedis.String(),
		Keterangan:  darah.Keterangan,
		Kadaluwarsa: helper.FormatTime(darah.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *DarahUseCase) GetByIdMedis(id string) model.DarahResponse {
	darah, err := u.Repository.FindByIdMedis(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Darah not found",
		})
	}

	response := model.DarahResponse{
		Id:          darah.Id.String(),
		IdMedis:     darah.IdMedis.String(),
		Keterangan:  darah.Keterangan,
		Kadaluwarsa: helper.FormatTime(darah.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *DarahUseCase) Update(request *model.DarahRequest, id, user string) model.DarahResponse {
	darah, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Darah not found",
		})
	}

	darah.IdMedis = helper.MustParse(request.IdMedis)
	darah.Keterangan = request.Keterangan
	darah.Kadaluwarsa = helper.ParseTime(request.Kadaluwarsa, "2006-01-02")
	darah.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&darah); err != nil {
		exception.PanicIfError(err, "Failed to update darah")
	}

	response := model.DarahResponse{
		Id:          darah.Id.String(),
		IdMedis:     darah.IdMedis.String(),
		Keterangan:  darah.Keterangan,
		Kadaluwarsa: helper.FormatTime(darah.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *DarahUseCase) Delete(id, user string) {
	darah, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Darah not found",
		})
	}

	darah.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&darah); err != nil {
		exception.PanicIfError(err, "Failed to delete darah")
	}
}
