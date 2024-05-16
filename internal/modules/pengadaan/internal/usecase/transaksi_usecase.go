package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/repository"
)

type TransaksiUseCase struct {
	Repository repository.TransaksiRepository
}

func NewTransaksiUseCase(repository *repository.TransaksiRepository) *TransaksiUseCase {
	return &TransaksiUseCase{
		Repository: *repository,
	}
}

func (u *TransaksiUseCase) Create(request *model.TransaksiRequest, user string) model.TransaksiResponse {
	updater := helper.MustParse(user)
	transaksi := entity.Transaksi{
		Id:      helper.MustNew(),
		IdStok:  helper.MustParse(request.IdStok),
		IdMedis: helper.MustParse(request.IdMedis),
		Batch:   request.Batch,
		Faktur:  request.Faktur,
		Jumlah:  request.Jumlah,
		Updater: updater,
	}

	if err := u.Repository.Insert(&transaksi); err != nil {
		exception.PanicIfError(err, "Failed to insert transaksi")
	}

	response := model.TransaksiResponse{
		Id:      transaksi.Id.String(),
		IdStok:  transaksi.IdStok.String(),
		IdMedis: transaksi.IdMedis.String(),
		Batch:   transaksi.Batch,
		Faktur:  transaksi.Faktur,
		Jumlah:  transaksi.Jumlah,
	}

	return response
}

func (u *TransaksiUseCase) Get() []model.TransaksiResponse {
	transaksi, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all transaksi")

	response := make([]model.TransaksiResponse, len(transaksi))
	for i, transaksi := range transaksi {
		response[i] = model.TransaksiResponse{
			Id:      transaksi.Id.String(),
			IdStok:  transaksi.IdStok.String(),
			IdMedis: transaksi.IdMedis.String(),
			Batch:   transaksi.Batch,
			Faktur:  transaksi.Faktur,
			Jumlah:  transaksi.Jumlah,
		}
	}

	return response
}

func (u *TransaksiUseCase) GetPage(page, size int) model.TransaksiPageResponse {
	transaksi, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged transaksi")

	response := make([]model.TransaksiResponse, len(transaksi))
	for i, transaksi := range transaksi {
		response[i] = model.TransaksiResponse{
			Id:      transaksi.Id.String(),
			IdStok:  transaksi.IdStok.String(),
			IdMedis: transaksi.IdMedis.String(),
			Batch:   transaksi.Batch,
			Faktur:  transaksi.Faktur,
			Jumlah:  transaksi.Jumlah,
		}
	}

	pagedResponse := model.TransaksiPageResponse{
		Page:      page,
		Size:      size,
		Total:     total,
		Transaksi: response,
	}

	return pagedResponse
}

func (u *TransaksiUseCase) GetById(id string) model.TransaksiResponse {
	transaksi, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Transaksi not found",
		})
	}

	response := model.TransaksiResponse{
		Id:      transaksi.Id.String(),
		IdStok:  transaksi.IdStok.String(),
		IdMedis: transaksi.IdMedis.String(),
		Batch:   transaksi.Batch,
		Faktur:  transaksi.Faktur,
		Jumlah:  transaksi.Jumlah,
	}

	return response
}

func (u *TransaksiUseCase) Update(request *model.TransaksiRequest, id, user string) model.TransaksiResponse {
	transaksi, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Transaksi not found",
		})
	}

	transaksi.IdStok = helper.MustParse(request.IdStok)
	transaksi.IdMedis = helper.MustParse(request.IdMedis)
	transaksi.Batch = request.Batch
	transaksi.Faktur = request.Faktur
	transaksi.Jumlah = request.Jumlah
	transaksi.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&transaksi); err != nil {
		exception.PanicIfError(err, "Failed to update transaksi")
	}

	response := model.TransaksiResponse{
		Id:      transaksi.Id.String(),
		IdStok:  transaksi.IdStok.String(),
		IdMedis: transaksi.IdMedis.String(),
		Batch:   transaksi.Batch,
		Faktur:  transaksi.Faktur,
		Jumlah:  transaksi.Jumlah,
	}

	return response
}

func (u *TransaksiUseCase) Delete(id, user string) {
	transaksi, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Transaksi not found",
		})
	}

	transaksi.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&transaksi); err != nil {
		exception.PanicIfError(err, "Failed to delete transaksi")
	}
}
