package service

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/helper"
	"github.com/fathoor/simkes-api/module/akun/akun/entity"
	"github.com/fathoor/simkes-api/module/akun/akun/model"
	"github.com/fathoor/simkes-api/module/akun/akun/repository"
	"github.com/fathoor/simkes-api/module/akun/akun/validation"
)

type akunServiceImpl struct {
	repository.AkunRepository
}

func (service *akunServiceImpl) Create(request *model.AkunRequest) error {
	valid := validation.ValidateAkunRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfError(err)

	akun := entity.Akun{
		NIP:      request.NIP,
		Email:    request.Email,
		Password: string(encrypted),
		RoleName: request.RoleName,
	}

	return service.AkunRepository.Insert(&akun)
}

func (service *akunServiceImpl) GetAll() ([]model.AkunResponse, error) {
	akun, err := service.AkunRepository.FindAll()

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			NIP:      akun.NIP,
			Email:    akun.Email,
			RoleName: akun.RoleName,
		}
	}

	return response, err
}

func (service *akunServiceImpl) GetByNIP(nip string) (model.AkunResponse, error) {
	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleName: akun.RoleName,
	}

	return response, nil
}

func (service *akunServiceImpl) PegawaiGetByNIP(nip string) (model.AkunResponse, error) {
	if nip == "Admin" {
		panic(exception.ForbiddenError{
			Message: "You are not allowed to access this resource",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleName: akun.RoleName,
	}

	return response, nil
}

func (service *akunServiceImpl) Update(nip string, request *model.AkunRequest) (model.AkunResponse, error) {
	valid := validation.ValidateAkunRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfError(err)

	if akun != (entity.Akun{}) {
		akun.NIP = request.NIP
		akun.Email = request.Email
		akun.Password = string(encrypted)
		akun.RoleName = request.RoleName
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleName: akun.RoleName,
	}

	return response, service.AkunRepository.Update(&akun)
}

func (service *akunServiceImpl) PegawaiUpdate(nip string, request *model.AkunUpdateRequest) (model.AkunResponse, error) {
	if nip == "Admin" {
		panic(exception.ForbiddenError{
			Message: "You are not allowed to access this resource",
		})
	}

	valid := validation.ValidateAkunUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfError(err)

	if akun != (entity.Akun{}) {
		akun.Email = request.Email
		akun.Password = string(encrypted)
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleName: akun.RoleName,
	}

	return response, service.AkunRepository.Update(&akun)
}

func (service *akunServiceImpl) Delete(nip string) error {
	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	return service.AkunRepository.Delete(&akun)
}

func ProvideAkunService(repository *repository.AkunRepository) AkunService {
	return &akunServiceImpl{*repository}
}
