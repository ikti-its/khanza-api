package service

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/jadwal-pegawai/entity"
	"github.com/fathoor/simkes-api/internal/jadwal-pegawai/model"
	"github.com/fathoor/simkes-api/internal/jadwal-pegawai/repository"
	"github.com/fathoor/simkes-api/internal/jadwal-pegawai/validation"
)

type jadwalPegawaiServiceImpl struct {
	repository.JadwalPegawaiRepository
}

func (service *jadwalPegawaiServiceImpl) Create(request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse {
	if valid := validation.ValidateJadwalPegawaiRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPegawai := entity.JadwalPegawai{
		NIP:       request.NIP,
		Tahun:     request.Tahun,
		Bulan:     request.Bulan,
		Hari:      request.Hari,
		ShiftNama: request.ShiftNama,
	}

	if err := service.JadwalPegawaiRepository.Insert(&jadwalPegawai); err != nil {
		exception.PanicIfError(err)
	}

	response := model.JadwalPegawaiResponse{
		NIP:       jadwalPegawai.NIP,
		Tahun:     jadwalPegawai.Tahun,
		Bulan:     jadwalPegawai.Bulan,
		Hari:      jadwalPegawai.Hari,
		ShiftNama: jadwalPegawai.ShiftNama,
	}

	return response
}

func (service *jadwalPegawaiServiceImpl) GetAll() []model.JadwalPegawaiResponse {
	jadwalPegawai, err := service.JadwalPegawaiRepository.FindAll()
	exception.PanicIfError(err)

	var response []model.JadwalPegawaiResponse
	for i, jadwalPegawai := range jadwalPegawai {
		response[i] = model.JadwalPegawaiResponse{
			NIP:       jadwalPegawai.NIP,
			Tahun:     jadwalPegawai.Tahun,
			Bulan:     jadwalPegawai.Bulan,
			Hari:      jadwalPegawai.Hari,
			ShiftNama: jadwalPegawai.ShiftNama,
		}
	}

	return response
}

func (service *jadwalPegawaiServiceImpl) GetByNIP(nip string) []model.JadwalPegawaiResponse {
	jadwalPegawai, err := service.JadwalPegawaiRepository.FindByNIP(nip)
	exception.PanicIfError(err)

	var response []model.JadwalPegawaiResponse
	for i, jadwalPegawai := range jadwalPegawai {
		response[i] = model.JadwalPegawaiResponse{
			NIP:       jadwalPegawai.NIP,
			Tahun:     jadwalPegawai.Tahun,
			Bulan:     jadwalPegawai.Bulan,
			Hari:      jadwalPegawai.Hari,
			ShiftNama: jadwalPegawai.ShiftNama,
		}
	}

	return response
}

func (service *jadwalPegawaiServiceImpl) GetByTahunBulan(tahun, bulan int16) []model.JadwalPegawaiResponse {
	jadwalPegawai, err := service.JadwalPegawaiRepository.FindByTahunBulan(tahun, bulan)
	exception.PanicIfError(err)

	var response []model.JadwalPegawaiResponse
	for i, jadwalPegawai := range jadwalPegawai {
		response[i] = model.JadwalPegawaiResponse{
			NIP:       jadwalPegawai.NIP,
			Tahun:     jadwalPegawai.Tahun,
			Bulan:     jadwalPegawai.Bulan,
			Hari:      jadwalPegawai.Hari,
			ShiftNama: jadwalPegawai.ShiftNama,
		}
	}

	return response
}

func (service *jadwalPegawaiServiceImpl) GetByPK(nip string, tahun, bulan, hari int16) model.JadwalPegawaiResponse {
	jadwalPegawai, err := service.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	response := model.JadwalPegawaiResponse{
		NIP:       jadwalPegawai.NIP,
		Tahun:     jadwalPegawai.Tahun,
		Bulan:     jadwalPegawai.Bulan,
		Hari:      jadwalPegawai.Hari,
		ShiftNama: jadwalPegawai.ShiftNama,
	}

	return response
}

func (service *jadwalPegawaiServiceImpl) Update(nip string, tahun, bulan, hari int16, request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse {
	if valid := validation.ValidateJadwalPegawaiRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPegawai, err := service.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	jadwalPegawai.ShiftNama = request.ShiftNama

	if err := service.JadwalPegawaiRepository.Update(&jadwalPegawai); err != nil {
		exception.PanicIfError(err)
	}

	response := model.JadwalPegawaiResponse{
		NIP:       jadwalPegawai.NIP,
		Tahun:     jadwalPegawai.Tahun,
		Bulan:     jadwalPegawai.Bulan,
		Hari:      jadwalPegawai.Hari,
		ShiftNama: jadwalPegawai.ShiftNama,
	}

	return response
}

func (service *jadwalPegawaiServiceImpl) Delete(nip string, tahun, bulan, hari int16) {
	jadwalPegawai, err := service.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	if err := service.JadwalPegawaiRepository.Delete(&jadwalPegawai); err != nil {
		exception.PanicIfError(err)
	}
}

func NewJadwalPegawaiServiceProvider(repository *repository.JadwalPegawaiRepository) JadwalPegawaiService {
	return &jadwalPegawaiServiceImpl{*repository}
}
