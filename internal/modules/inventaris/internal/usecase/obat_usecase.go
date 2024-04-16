package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type ObatUseCase struct {
	Repository repository.ObatRepository
}

func NewObatUseCase(repository *repository.ObatRepository) *ObatUseCase {
	return &ObatUseCase{
		Repository: *repository,
	}
}

func (u *ObatUseCase) Create(request *model.ObatRequest, user string) model.ObatResponse {
	updater := helper.MustParse(user)
	obat := entity.Obat{
		Id:          helper.MustNew(),
		IdMedis:     helper.MustParse(request.IdMedis),
		Industri:    request.Industri,
		Kandungan:   request.Kandungan,
		SatuanBesar: request.SatuanBesar,
		SatuanKecil: request.SatuanKecil,
		Isi:         request.Isi,
		Kapasitas:   request.Kapasitas,
		Jenis:       request.Jenis,
		Kategori:    request.Kategori,
		Golongan:    request.Golongan,
		Kadaluwarsa: helper.ParseTime(request.Kadaluwarsa, "2006-01-02"),
		Updater:     updater,
	}

	if err := u.Repository.Insert(&obat); err != nil {
		exception.PanicIfError(err, "Failed to insert obat")
	}

	response := model.ObatResponse{
		Id:          obat.Id.String(),
		IdMedis:     obat.IdMedis.String(),
		Industri:    obat.Industri,
		Kandungan:   obat.Kandungan,
		SatuanBesar: obat.SatuanBesar,
		SatuanKecil: obat.SatuanKecil,
		Isi:         obat.Isi,
		Kapasitas:   obat.Kapasitas,
		Jenis:       obat.Jenis,
		Kategori:    obat.Kategori,
		Golongan:    obat.Golongan,
		Kadaluwarsa: helper.FormatTime(obat.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *ObatUseCase) Get() []model.ObatResponse {
	obat, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all obat")

	response := make([]model.ObatResponse, len(obat))
	for i, obat := range obat {
		response[i] = model.ObatResponse{
			Id:          obat.Id.String(),
			IdMedis:     obat.IdMedis.String(),
			Industri:    obat.Industri,
			Kandungan:   obat.Kandungan,
			SatuanBesar: obat.SatuanBesar,
			SatuanKecil: obat.SatuanKecil,
			Isi:         obat.Isi,
			Kapasitas:   obat.Kapasitas,
			Jenis:       obat.Jenis,
			Kategori:    obat.Kategori,
			Golongan:    obat.Golongan,
			Kadaluwarsa: helper.FormatTime(obat.Kadaluwarsa, "2006-01-02"),
		}
	}

	return response
}

func (u *ObatUseCase) GetPage(page, size int) model.ObatPageResponse {
	obat, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged obat")

	response := make([]model.ObatResponse, len(obat))
	for i, obat := range obat {
		response[i] = model.ObatResponse{
			Id:          obat.Id.String(),
			IdMedis:     obat.IdMedis.String(),
			Industri:    obat.Industri,
			Kandungan:   obat.Kandungan,
			SatuanBesar: obat.SatuanBesar,
			SatuanKecil: obat.SatuanKecil,
			Isi:         obat.Isi,
			Kapasitas:   obat.Kapasitas,
			Jenis:       obat.Jenis,
			Kategori:    obat.Kategori,
			Golongan:    obat.Golongan,
			Kadaluwarsa: helper.FormatTime(obat.Kadaluwarsa, "2006-01-02"),
		}
	}

	pagedResponse := model.ObatPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Obat:  response,
	}

	return pagedResponse
}

func (u *ObatUseCase) GetById(id string) model.ObatResponse {
	obat, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Obat not found",
		})
	}

	response := model.ObatResponse{
		Id:          obat.Id.String(),
		IdMedis:     obat.IdMedis.String(),
		Industri:    obat.Industri,
		Kandungan:   obat.Kandungan,
		SatuanBesar: obat.SatuanBesar,
		SatuanKecil: obat.SatuanKecil,
		Isi:         obat.Isi,
		Kapasitas:   obat.Kapasitas,
		Jenis:       obat.Jenis,
		Kategori:    obat.Kategori,
		Golongan:    obat.Golongan,
		Kadaluwarsa: helper.FormatTime(obat.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *ObatUseCase) Update(request *model.ObatRequest, id, user string) model.ObatResponse {
	obat, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Obat not found",
		})
	}

	obat.IdMedis = helper.MustParse(request.IdMedis)
	obat.Industri = request.Industri
	obat.Kandungan = request.Kandungan
	obat.SatuanBesar = request.SatuanBesar
	obat.SatuanKecil = request.SatuanKecil
	obat.Isi = request.Isi
	obat.Kapasitas = request.Kapasitas
	obat.Jenis = request.Jenis
	obat.Kategori = request.Kategori
	obat.Golongan = request.Golongan
	obat.Kadaluwarsa = helper.ParseTime(request.Kadaluwarsa, "2006-01-02")
	obat.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&obat); err != nil {
		exception.PanicIfError(err, "Failed to update obat")
	}

	response := model.ObatResponse{
		Id:          obat.Id.String(),
		IdMedis:     obat.IdMedis.String(),
		Industri:    obat.Industri,
		Kandungan:   obat.Kandungan,
		SatuanBesar: obat.SatuanBesar,
		SatuanKecil: obat.SatuanKecil,
		Isi:         obat.Isi,
		Kapasitas:   obat.Kapasitas,
		Jenis:       obat.Jenis,
		Kategori:    obat.Kategori,
		Golongan:    obat.Golongan,
		Kadaluwarsa: helper.FormatTime(obat.Kadaluwarsa, "2006-01-02"),
	}

	return response
}

func (u *ObatUseCase) Delete(id, user string) {
	obat, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Obat not found",
		})
	}

	obat.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&obat); err != nil {
		exception.PanicIfError(err, "Failed to delete obat")
	}
}
