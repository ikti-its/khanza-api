package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type MedisUseCase struct {
	Repository repository.MedisRepository
}

func NewMedisUseCase(repository *repository.MedisRepository) *MedisUseCase {
	return &MedisUseCase{
		Repository: *repository,
	}
}

func (u *MedisUseCase) Create(request *model.MedisRequest, user string) model.MedisResponse {
	updater := helper.MustParse(user)
	medis := entity.Medis{
		Id:      helper.MustNew(),
		Nama:    request.Nama,
		Jenis:   request.Jenis,
		Satuan:  request.Satuan,
		Harga:   request.Harga,
		Stok:    request.Stok,
		Updater: updater,
	}

	if err := u.Repository.Insert(&medis); err != nil {
		exception.PanicIfError(err, "Failed to insert barang medis")
	}

	response := model.MedisResponse{
		Id:     medis.Id.String(),
		Nama:   medis.Nama,
		Jenis:  medis.Jenis,
		Satuan: medis.Satuan,
		Harga:  medis.Harga,
		Stok:   medis.Stok,
	}

	return response
}

func (u *MedisUseCase) Get() []model.MedisResponse {
	medis, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all barang medis")

	response := make([]model.MedisResponse, len(medis))
	for i, medis := range medis {
		response[i] = model.MedisResponse{
			Id:     medis.Id.String(),
			Nama:   medis.Nama,
			Jenis:  medis.Jenis,
			Satuan: medis.Satuan,
			Harga:  medis.Harga,
			Stok:   medis.Stok,
		}
	}

	return response
}

func (u *MedisUseCase) GetPage(page, size int) model.MedisPageResponse {
	medis, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged barang medis")

	response := make([]model.MedisResponse, len(medis))
	for i, medis := range medis {
		response[i] = model.MedisResponse{
			Id:     medis.Id.String(),
			Nama:   medis.Nama,
			Jenis:  medis.Jenis,
			Satuan: medis.Satuan,
			Harga:  medis.Harga,
			Stok:   medis.Stok,
		}
	}

	pagedResponse := model.MedisPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Medis: response,
	}

	return pagedResponse
}

func (u *MedisUseCase) GetByJenis(jenis string) []model.MedisResponse {
	medis, err := u.Repository.FindByJenis(jenis)
	exception.PanicIfError(err, "Failed to get barang medis")

	response := make([]model.MedisResponse, len(medis))
	for i, medis := range medis {
		response[i] = model.MedisResponse{
			Id:     medis.Id.String(),
			Nama:   medis.Nama,
			Jenis:  medis.Jenis,
			Satuan: medis.Satuan,
			Harga:  medis.Harga,
			Stok:   medis.Stok,
		}
	}

	return response
}

func (u *MedisUseCase) GetById(id string) model.MedisResponse {
	medis, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Barang medis not found",
		})
	}

	response := model.MedisResponse{
		Id:     medis.Id.String(),
		Nama:   medis.Nama,
		Jenis:  medis.Jenis,
		Satuan: medis.Satuan,
		Harga:  medis.Harga,
		Stok:   medis.Stok,
	}

	return response
}

func (u *MedisUseCase) Update(request *model.MedisRequest, id, user string) model.MedisResponse {
	medis, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Barang medis not found",
		})
	}

	medis.Nama = request.Nama
	medis.Jenis = request.Jenis
	medis.Satuan = request.Satuan
	medis.Harga = request.Harga
	medis.Stok = request.Stok
	medis.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&medis); err != nil {
		exception.PanicIfError(err, "Failed to update barang medis")
	}

	response := model.MedisResponse{
		Id:     medis.Id.String(),
		Nama:   medis.Nama,
		Jenis:  medis.Jenis,
		Satuan: medis.Satuan,
		Harga:  medis.Harga,
		Stok:   medis.Stok,
	}

	return response
}

func (u *MedisUseCase) Delete(id, user string) {
	medis, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Barang medis not found",
		})
	}

	medis.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&medis); err != nil {
		exception.PanicIfError(err, "Failed to delete barang medis")
	}
}
