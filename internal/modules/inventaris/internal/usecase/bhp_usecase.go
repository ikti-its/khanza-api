package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type BhpUseCase struct {
	Repository repository.BhpRepository
}

func NewBhpUseCase(repository *repository.BhpRepository) *BhpUseCase {
	return &BhpUseCase{
		Repository: *repository,
	}
}

func (u *BhpUseCase) Create(request *model.BhpRequest, user string) model.BhpResponse {
	updater := helper.MustParse(user)
	bhp := entity.Bhp{
		Id:          helper.MustNew(),
		IdMedis:     helper.MustParse(request.IdMedis),
		Satuan:      request.Satuan,
		Jumlah:      request.Jumlah,
		Kadaluwarsa: helper.ParseTime(request.Kadaluwarsa, "2006-01-02"),
		Updater:     updater,
	}

	if err := u.Repository.Insert(&bhp); err != nil {
		exception.PanicIfError(err, "Failed to insert bhp")
	}

	response := model.BhpResponse{
		Id:          bhp.Id.String(),
		IdMedis:     bhp.IdMedis.String(),
		Satuan:      bhp.Satuan,
		Jumlah:      bhp.Jumlah,
		Kadaluwarsa: helper.FormatTime(bhp.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *BhpUseCase) Get() []model.BhpResponse {
	bhp, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all bhp")

	response := make([]model.BhpResponse, len(bhp))
	for i, bhp := range bhp {
		response[i] = model.BhpResponse{
			Id:          bhp.Id.String(),
			IdMedis:     bhp.IdMedis.String(),
			Satuan:      bhp.Satuan,
			Jumlah:      bhp.Jumlah,
			Kadaluwarsa: helper.FormatTime(bhp.Kadaluwarsa, "2006-01-02"),
		}
	}

	return response
}

func (u *BhpUseCase) GetPage(page, size int) model.BhpPageResponse {
	bhp, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged bhp")

	response := make([]model.BhpResponse, len(bhp))
	for i, bhp := range bhp {
		response[i] = model.BhpResponse{
			Id:          bhp.Id.String(),
			IdMedis:     bhp.IdMedis.String(),
			Satuan:      bhp.Satuan,
			Jumlah:      bhp.Jumlah,
			Kadaluwarsa: helper.FormatTime(bhp.Kadaluwarsa, "2006-01-02"),
		}
	}

	pagedResponse := model.BhpPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Bhp:   response,
	}

	return pagedResponse
}

func (u *BhpUseCase) GetById(id string) model.BhpResponse {
	bhp, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Bhp not found",
		})
	}

	response := model.BhpResponse{
		Id:          bhp.Id.String(),
		IdMedis:     bhp.IdMedis.String(),
		Satuan:      bhp.Satuan,
		Jumlah:      bhp.Jumlah,
		Kadaluwarsa: helper.FormatTime(bhp.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *BhpUseCase) Update(request *model.BhpRequest, id, user string) model.BhpResponse {
	bhp, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Bhp not found",
		})
	}

	bhp.IdMedis = helper.MustParse(request.IdMedis)
	bhp.Satuan = request.Satuan
	bhp.Jumlah = request.Jumlah
	bhp.Kadaluwarsa = helper.ParseTime(request.Kadaluwarsa, "2006-01-02")
	bhp.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&bhp); err != nil {
		exception.PanicIfError(err, "Failed to update bhp")
	}

	response := model.BhpResponse{
		Id:          bhp.Id.String(),
		IdMedis:     bhp.IdMedis.String(),
		Satuan:      bhp.Satuan,
		Jumlah:      bhp.Jumlah,
		Kadaluwarsa: helper.FormatTime(bhp.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *BhpUseCase) Delete(id, user string) {
	bhp, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Bhp not found",
		})
	}

	bhp.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&bhp); err != nil {
		exception.PanicIfError(err, "Failed to delete bhp")
	}
}
