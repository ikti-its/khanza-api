package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/repository"
)

type OrganisasiUseCase struct {
	Repository repository.OrganisasiRepository
}

func NewOrganisasiUseCase(repository *repository.OrganisasiRepository) *OrganisasiUseCase {
	return &OrganisasiUseCase{
		Repository: *repository,
	}
}

func (u *OrganisasiUseCase) Create(request *model.OrganisasiRequest, user string) model.OrganisasiResponse {
	updater := helper.MustParse(user)
	organisasi := entity.Organisasi{
		Id:        helper.MustNew(),
		Nama:      request.Nama,
		Alamat:    request.Alamat,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Radius:    request.Radius,
		Updater:   updater,
	}

	if err := u.Repository.Insert(&organisasi); err != nil {
		exception.PanicIfError(err, "Failed to insert organisasi")
	}

	response := model.OrganisasiResponse{
		Id:        organisasi.Id.String(),
		Nama:      organisasi.Nama,
		Alamat:    organisasi.Alamat,
		Latitude:  organisasi.Latitude,
		Longitude: organisasi.Longitude,
		Radius:    organisasi.Radius,
	}

	return response
}

func (u *OrganisasiUseCase) Get() []model.OrganisasiResponse {
	organisasi, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all organisasi")

	response := make([]model.OrganisasiResponse, len(organisasi))
	for i, organisasi := range organisasi {
		response[i] = model.OrganisasiResponse{
			Id:        organisasi.Id.String(),
			Nama:      organisasi.Nama,
			Alamat:    organisasi.Alamat,
			Latitude:  organisasi.Latitude,
			Longitude: organisasi.Longitude,
			Radius:    organisasi.Radius,
		}
	}

	return response
}

func (u *OrganisasiUseCase) GetPage(page, size int) model.OrganisasiPageResponse {
	organisasi, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged organisasi")

	response := make([]model.OrganisasiResponse, len(organisasi))
	for i, organisasi := range organisasi {
		response[i] = model.OrganisasiResponse{
			Id:        organisasi.Id.String(),
			Nama:      organisasi.Nama,
			Alamat:    organisasi.Alamat,
			Latitude:  organisasi.Latitude,
			Longitude: organisasi.Longitude,
			Radius:    organisasi.Radius,
		}
	}

	pagedResponse := model.OrganisasiPageResponse{
		Page:       page,
		Size:       size,
		Total:      total,
		Organisasi: response,
	}

	return pagedResponse
}

func (u *OrganisasiUseCase) GetCurrent() model.OrganisasiResponse {
	organisasi, err := u.Repository.FindCurrent()
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Organisasi not found",
		})
	}

	response := model.OrganisasiResponse{
		Id:        organisasi.Id.String(),
		Nama:      organisasi.Nama,
		Alamat:    organisasi.Alamat,
		Latitude:  organisasi.Latitude,
		Longitude: organisasi.Longitude,
		Radius:    organisasi.Radius,
	}

	return response
}

func (u *OrganisasiUseCase) GetById(id string) model.OrganisasiResponse {
	organisasi, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Organisasi not found",
		})
	}

	response := model.OrganisasiResponse{
		Id:        organisasi.Id.String(),
		Nama:      organisasi.Nama,
		Alamat:    organisasi.Alamat,
		Latitude:  organisasi.Latitude,
		Longitude: organisasi.Longitude,
		Radius:    organisasi.Radius,
	}

	return response
}

func (u *OrganisasiUseCase) Update(request *model.OrganisasiRequest, id, user string) model.OrganisasiResponse {
	organisasi, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Organisasi not found",
		})
	}

	organisasi.Nama = request.Nama
	organisasi.Alamat = request.Alamat
	organisasi.Latitude = request.Latitude
	organisasi.Longitude = request.Longitude
	organisasi.Radius = request.Radius
	organisasi.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&organisasi); err != nil {
		exception.PanicIfError(err, "Failed to update organisasi")
	}

	response := model.OrganisasiResponse{
		Id:        organisasi.Id.String(),
		Nama:      organisasi.Nama,
		Alamat:    organisasi.Alamat,
		Latitude:  organisasi.Latitude,
		Longitude: organisasi.Longitude,
		Radius:    organisasi.Radius,
	}

	return response
}

func (u *OrganisasiUseCase) Delete(id, user string) {
	organisasi, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Organisasi not found",
		})
	}

	organisasi.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&organisasi); err != nil {
		exception.PanicIfError(err, "Failed to delete organisasi")
	}
}
