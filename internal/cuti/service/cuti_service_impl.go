package service

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/cuti/entity"
	"github.com/fathoor/simkes-api/internal/cuti/model"
	"github.com/fathoor/simkes-api/internal/cuti/repository"
	"github.com/fathoor/simkes-api/internal/cuti/validation"
	"github.com/google/uuid"
	"time"
)

type cutiServiceImpl struct {
	repository.CutiRepository
}

func (service *cutiServiceImpl) Create(request *model.CutiCreateRequest) model.CutiResponse {
	if valid := validation.ValidateCutiCreateRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	tanggalMulai, err := time.Parse("2006-01-02", request.TanggalMulai)
	exception.PanicIfError(err)

	tanggalSelesai, err := time.Parse("2006-01-02", request.TanggalSelesai)
	exception.PanicIfError(err)

	cuti := entity.Cuti{
		ID:             uuid.New(),
		NIP:            request.NIP,
		TanggalMulai:   tanggalMulai,
		TanggalSelesai: tanggalSelesai,
		Keterangan:     request.Keterangan,
	}

	if err := service.CutiRepository.Insert(&cuti); err != nil {
		exception.PanicIfError(err)
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (service *cutiServiceImpl) GetAll() []model.CutiResponse {
	cuti, err := service.CutiRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.CutiResponse, len(cuti))
	for i, cuti := range cuti {
		response[i] = model.CutiResponse{
			ID:             cuti.ID.String(),
			NIP:            cuti.NIP,
			TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
			TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
			Keterangan:     cuti.Keterangan,
			Status:         cuti.Status,
		}
	}

	return response
}

func (service *cutiServiceImpl) GetByNIP(nip string) []model.CutiResponse {
	cuti, err := service.CutiRepository.FindByNIP(nip)
	exception.PanicIfError(err)

	response := make([]model.CutiResponse, len(cuti))
	for i, cuti := range cuti {
		response[i] = model.CutiResponse{
			ID:             cuti.ID.String(),
			NIP:            cuti.NIP,
			TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
			TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
			Keterangan:     cuti.Keterangan,
			Status:         cuti.Status,
		}
	}

	return response
}

func (service *cutiServiceImpl) GetByID(id string) model.CutiResponse {
	cutiID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	cuti, err := service.CutiRepository.FindByID(cutiID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (service *cutiServiceImpl) Update(id string, request *model.CutiUpdateRequest) model.CutiResponse {
	if valid := validation.ValidateCutiUpdateRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	cutiID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	cuti, err := service.CutiRepository.FindByID(cutiID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	tanggalMulai, err := time.Parse("2006-01-02", request.TanggalMulai)
	exception.PanicIfError(err)

	tanggalSelesai, err := time.Parse("2006-01-02", request.TanggalSelesai)
	exception.PanicIfError(err)

	cuti.TanggalMulai = tanggalMulai
	cuti.TanggalSelesai = tanggalSelesai
	cuti.Keterangan = request.Keterangan

	if err := service.CutiRepository.Update(&cuti); err != nil {
		exception.PanicIfError(err)
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (service *cutiServiceImpl) UpdateStatus(id string, request *model.CutiUpdateRequest) model.CutiResponse {
	if valid := validation.ValidateCutiUpdateRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	cutiID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	cuti, err := service.CutiRepository.FindByID(cutiID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	tanggalMulai, err := time.Parse("2006-01-02", request.TanggalMulai)
	exception.PanicIfError(err)

	tanggalSelesai, err := time.Parse("2006-01-02", request.TanggalSelesai)
	exception.PanicIfError(err)

	cuti.TanggalMulai = tanggalMulai
	cuti.TanggalSelesai = tanggalSelesai
	cuti.Keterangan = request.Keterangan
	cuti.Status = request.Status

	if err := service.CutiRepository.Update(&cuti); err != nil {
		exception.PanicIfError(err)
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (service *cutiServiceImpl) Delete(id string) {
	cutiID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	cuti, err := service.CutiRepository.FindByID(cutiID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	if err := service.CutiRepository.Delete(&cuti); err != nil {
		exception.PanicIfError(err)
	}
}

func NewCutiServiceProvider(repository *repository.CutiRepository) CutiService {
	return &cutiServiceImpl{*repository}
}
