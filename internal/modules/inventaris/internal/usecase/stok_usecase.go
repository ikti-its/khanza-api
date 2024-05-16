package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type StokUseCase struct {
	Repository repository.StokRepository
}

func NewStokUseCase(repository *repository.StokRepository) *StokUseCase {
	return &StokUseCase{
		Repository: *repository,
	}
}

func (u *StokUseCase) Create(request *model.StokRequest, user string) model.StokResponse {
	updater := helper.MustParse(user)
	stok := entity.Stok{
		Id:         helper.MustNew(),
		Nomor:      request.Nomor,
		IdPegawai:  helper.MustParse(request.IdPegawai),
		Tanggal:    helper.ParseTime(request.Tanggal, "2006-01-02"),
		Keterangan: request.Keterangan,
		Updater:    updater,
	}

	if err := u.Repository.Insert(&stok); err != nil {
		exception.PanicIfError(err, "Failed to insert stok")
	}

	response := model.StokResponse{
		Id:         stok.Id.String(),
		Nomor:      stok.Nomor,
		IdPegawai:  stok.IdPegawai.String(),
		Tanggal:    helper.FormatTime(stok.Tanggal, "2006-01-02"),
		Keterangan: stok.Keterangan,
	}

	return response
}

func (u *StokUseCase) Get() []model.StokResponse {
	stok, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all stok")

	response := make([]model.StokResponse, len(stok))
	for i, stok := range stok {
		response[i] = model.StokResponse{
			Id:         stok.Id.String(),
			Nomor:      stok.Nomor,
			IdPegawai:  stok.IdPegawai.String(),
			Tanggal:    helper.FormatTime(stok.Tanggal, "2006-01-02"),
			Keterangan: stok.Keterangan,
		}
	}

	return response
}

func (u *StokUseCase) GetPage(page, size int) model.StokPageResponse {
	stok, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged stok")

	response := make([]model.StokResponse, len(stok))
	for i, stok := range stok {
		response[i] = model.StokResponse{
			Id:         stok.Id.String(),
			Nomor:      stok.Nomor,
			IdPegawai:  stok.IdPegawai.String(),
			Tanggal:    helper.FormatTime(stok.Tanggal, "2006-01-02"),
			Keterangan: stok.Keterangan,
		}
	}

	pagedResponse := model.StokPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Stok:  response,
	}

	return pagedResponse
}

func (u *StokUseCase) GetById(id string) model.StokResponse {
	stok, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Stok not found",
		})
	}

	response := model.StokResponse{
		Id:         stok.Id.String(),
		Nomor:      stok.Nomor,
		IdPegawai:  stok.IdPegawai.String(),
		Tanggal:    helper.FormatTime(stok.Tanggal, "2006-01-02"),
		Keterangan: stok.Keterangan,
	}

	return response
}

func (u *StokUseCase) Update(request *model.StokRequest, id, user string) model.StokResponse {
	stok, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Stok not found",
		})
	}

	stok.Nomor = request.Nomor
	stok.IdPegawai = helper.MustParse(request.IdPegawai)
	stok.Tanggal = helper.ParseTime(request.Tanggal, "2006-01-02")
	stok.Keterangan = request.Keterangan
	stok.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&stok); err != nil {
		exception.PanicIfError(err, "Failed to update stok")
	}

	response := model.StokResponse{
		Id:         stok.Id.String(),
		Nomor:      stok.Nomor,
		IdPegawai:  stok.IdPegawai.String(),
		Tanggal:    helper.FormatTime(stok.Tanggal, "2006-01-02"),
		Keterangan: stok.Keterangan,
	}

	return response
}

func (u *StokUseCase) Delete(id, user string) {
	stok, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Stok not found",
		})
	}

	stok.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&stok); err != nil {
		exception.PanicIfError(err, "Failed to delete stok")
	}
}
