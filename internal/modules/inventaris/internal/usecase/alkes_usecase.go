package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type AlkesUseCase struct {
	Repository repository.AlkesRepository
}

func NewAlkesUseCase(repository *repository.AlkesRepository) *AlkesUseCase {
	return &AlkesUseCase{
		Repository: *repository,
	}
}

func (u *AlkesUseCase) Create(request *model.AlkesRequest, user string) model.AlkesResponse {
	updater := helper.MustParse(user)
	alkes := entity.Alkes{
		Id:      helper.MustNew(),
		IdMedis: helper.MustParse(request.IdMedis),
		Merek:   request.Merek,
		Updater: updater,
	}

	if err := u.Repository.Insert(&alkes); err != nil {
		exception.PanicIfError(err, "Failed to insert alkes")
	}

	response := model.AlkesResponse{
		Id:      alkes.Id.String(),
		IdMedis: alkes.IdMedis.String(),
		Merek:   alkes.Merek,
	}

	return response
}

func (u *AlkesUseCase) Get() []model.AlkesResponse {
	alkes, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all alkes")

	response := make([]model.AlkesResponse, len(alkes))
	for i, alkes := range alkes {
		response[i] = model.AlkesResponse{
			Id:      alkes.Id.String(),
			IdMedis: alkes.IdMedis.String(),
			Merek:   alkes.Merek,
		}
	}

	return response
}

func (u *AlkesUseCase) GetPage(page, size int) model.AlkesPageResponse {
	alkes, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged alkes")

	response := make([]model.AlkesResponse, len(alkes))
	for i, alkes := range alkes {
		response[i] = model.AlkesResponse{
			Id:      alkes.Id.String(),
			IdMedis: alkes.IdMedis.String(),
			Merek:   alkes.Merek,
		}
	}

	pagedResponse := model.AlkesPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Alkes: response,
	}

	return pagedResponse
}

func (u *AlkesUseCase) GetById(id string) model.AlkesResponse {
	alkes, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alkes not found",
		})
	}

	response := model.AlkesResponse{
		Id:      alkes.Id.String(),
		IdMedis: alkes.IdMedis.String(),
		Merek:   alkes.Merek,
	}

	return response
}

func (u *AlkesUseCase) GetByIdMedis(id string) model.AlkesResponse {
	alkes, err := u.Repository.FindByIdMedis(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alkes not found",
		})
	}

	response := model.AlkesResponse{
		Id:      alkes.Id.String(),
		IdMedis: alkes.IdMedis.String(),
		Merek:   alkes.Merek,
	}

	return response
}

func (u *AlkesUseCase) Update(request *model.AlkesRequest, id, user string) model.AlkesResponse {
	alkes, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alkes not found",
		})
	}

	alkes.IdMedis = helper.MustParse(request.IdMedis)
	alkes.Merek = request.Merek
	alkes.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&alkes); err != nil {
		exception.PanicIfError(err, "Failed to update alkes")
	}

	response := model.AlkesResponse{
		Id:      alkes.Id.String(),
		IdMedis: alkes.IdMedis.String(),
		Merek:   alkes.Merek,
	}

	return response
}

func (u *AlkesUseCase) Delete(id, user string) {
	alkes, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Alkes not found",
		})
	}

	alkes.Updater = helper.MustParse(user)

	if err := u.Repository.Delete(&alkes); err != nil {
		exception.PanicIfError(err, "Failed to delete alkes")
	}
}
