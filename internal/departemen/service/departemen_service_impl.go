package service

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/departemen/entity"
	"github.com/fathoor/simkes-api/internal/departemen/model"
	"github.com/fathoor/simkes-api/internal/departemen/repository"
	"github.com/fathoor/simkes-api/internal/departemen/validation"
)

type departemenServiceImpl struct {
	repository.DepartemenRepository
}

func (service *departemenServiceImpl) Create(request *model.DepartemenRequest) model.DepartemenResponse {
	if err := validation.ValidateDepartemenRequest(request); err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	departemen := entity.Departemen{
		Nama: request.Nama,
	}

	if err := service.DepartemenRepository.Insert(&departemen); err != nil {
		exception.PanicIfError(err)
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (service *departemenServiceImpl) GetAll() []model.DepartemenResponse {
	departemen, err := service.DepartemenRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.DepartemenResponse, len(departemen))
	for i, departemen := range departemen {
		response[i] = model.DepartemenResponse{
			Nama: departemen.Nama,
		}
	}

	return response
}

func (service *departemenServiceImpl) GetByDepartemen(d string) model.DepartemenResponse {
	departemen, err := service.DepartemenRepository.FindByDepartemen(d)
	exception.PanicIfError(err)

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (service *departemenServiceImpl) Update(d string, request *model.DepartemenRequest) model.DepartemenResponse {
	if err := validation.ValidateDepartemenRequest(request); err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	if _, err := service.DepartemenRepository.FindByDepartemen(request.Nama); err == nil {
		panic(exception.BadRequestError{
			Message: "Departemen already exists",
		})
	}

	departemen, err := service.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	departemen.Nama = request.Nama

	if err := service.DepartemenRepository.Update(&departemen); err != nil {
		exception.PanicIfError(err)
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (service *departemenServiceImpl) Delete(d string) {
	departemen, err := service.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	if err := service.DepartemenRepository.Delete(&departemen); err != nil {
		exception.PanicIfError(err)
	}
}

func NewDepartemenServiceProvider(repository *repository.DepartemenRepository) DepartemenService {
	return &departemenServiceImpl{*repository}
}
