package service

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/jabatan/entity"
	"github.com/fathoor/simkes-api/internal/jabatan/model"
	"github.com/fathoor/simkes-api/internal/jabatan/repository"
	"github.com/fathoor/simkes-api/internal/jabatan/validation"
)

type jabatanServiceImpl struct {
	repository.JabatanRepository
}

func (service *jabatanServiceImpl) Create(request *model.JabatanRequest) model.JabatanResponse {
	if err := validation.ValidateJabatanRequest(request); err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jabatan := entity.Jabatan{
		Nama:      request.Nama,
		Jenjang:   request.Jenjang,
		GajiPokok: request.GajiPokok,
		Tunjangan: request.Tunjangan,
	}

	if err := service.JabatanRepository.Insert(&jabatan); err != nil {
		exception.PanicIfError(err)
	}

	response := model.JabatanResponse{
		Nama:      jabatan.Nama,
		Jenjang:   jabatan.Jenjang,
		GajiPokok: jabatan.GajiPokok,
		Tunjangan: jabatan.Tunjangan,
	}

	return response
}

func (service *jabatanServiceImpl) GetAll() []model.JabatanResponse {
	jabatan, err := service.JabatanRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.JabatanResponse, len(jabatan))
	for i, jabatan := range jabatan {
		response[i] = model.JabatanResponse{
			Nama:      jabatan.Nama,
			Jenjang:   jabatan.Jenjang,
			GajiPokok: jabatan.GajiPokok,
			Tunjangan: jabatan.Tunjangan,
		}
	}

	return response
}

func (service *jabatanServiceImpl) GetByJabatan(j string) model.JabatanResponse {
	jabatan, err := service.JabatanRepository.FindByJabatan(j)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	response := model.JabatanResponse{
		Nama:      jabatan.Nama,
		Jenjang:   jabatan.Jenjang,
		GajiPokok: jabatan.GajiPokok,
		Tunjangan: jabatan.Tunjangan,
	}

	return response
}

func (service *jabatanServiceImpl) Update(j string, request *model.JabatanRequest) model.JabatanResponse {
	if valid := validation.ValidateJabatanRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	if _, err := service.JabatanRepository.FindByJabatan(request.Nama); err == nil {
		panic(exception.BadRequestError{
			Message: "Jabatan already exists",
		})
	}

	jabatan, err := service.JabatanRepository.FindByJabatan(j)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	jabatan.Nama = request.Nama
	jabatan.Jenjang = request.Jenjang
	jabatan.GajiPokok = request.GajiPokok
	jabatan.Tunjangan = request.Tunjangan

	if err := service.JabatanRepository.Update(&jabatan); err != nil {
		exception.PanicIfError(err)
	}

	response := model.JabatanResponse{
		Nama:      jabatan.Nama,
		Jenjang:   jabatan.Jenjang,
		GajiPokok: jabatan.GajiPokok,
		Tunjangan: jabatan.Tunjangan,
	}

	return response
}

func (service *jabatanServiceImpl) Delete(j string) {
	jabatan, err := service.JabatanRepository.FindByJabatan(j)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	if err := service.JabatanRepository.Delete(&jabatan); err != nil {
		exception.PanicIfError(err)
	}
}

func NewJabatanServiceProvider(repository *repository.JabatanRepository) JabatanService {
	return &jabatanServiceImpl{*repository}
}
