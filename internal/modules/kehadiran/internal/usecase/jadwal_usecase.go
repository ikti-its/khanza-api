package usecase

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/helper"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/repository"
)

type JadwalUseCase struct {
	Repository repository.JadwalRepository
}

func NewJadwalUseCase(repository *repository.JadwalRepository) *JadwalUseCase {
	return &JadwalUseCase{
		Repository: *repository,
	}
}

func (u *JadwalUseCase) Get() []model.JadwalResponse {
	jadwal, err := u.Repository.Find()
	if err != nil {
		exception.PanicIfError(err, "Failed to get jadwal")
	}

	response := make([]model.JadwalResponse, len(jadwal))
	for i, jadwal := range jadwal {
		response[i] = model.JadwalResponse{
			Id:        jadwal.Id.String(),
			IdPegawai: jadwal.IdPegawai.String(),
			IdHari:    jadwal.IdHari,
			IdShift:   jadwal.IdShift,
			JamMasuk:  jadwal.JamMasuk,
			JamPulang: jadwal.JamPulang,
		}
	}

	return response
}

func (u *JadwalUseCase) GetByHariId(id int) []model.JadwalResponse {
	jadwal, err := u.Repository.FindByHariId(id)
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Jadwal not found",
		})
	}

	response := make([]model.JadwalResponse, len(jadwal))
	for i, jadwal := range jadwal {
		response[i] = model.JadwalResponse{
			Id:        jadwal.Id.String(),
			IdPegawai: jadwal.IdPegawai.String(),
			IdHari:    jadwal.IdHari,
			IdShift:   jadwal.IdShift,
			JamMasuk:  jadwal.JamMasuk,
			JamPulang: jadwal.JamPulang,
		}
	}

	return response
}

func (u *JadwalUseCase) GetByPegawaiId(id string) []model.JadwalResponse {
	jadwal, err := u.Repository.FindByPegawaiId(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Jadwal not found",
		})
	}

	response := make([]model.JadwalResponse, len(jadwal))
	for i, jadwal := range jadwal {
		response[i] = model.JadwalResponse{
			Id:        jadwal.Id.String(),
			IdPegawai: jadwal.IdPegawai.String(),
			IdHari:    jadwal.IdHari,
			IdShift:   jadwal.IdShift,
			JamMasuk:  jadwal.JamMasuk,
			JamPulang: jadwal.JamPulang,
		}
	}

	return response
}

func (u *JadwalUseCase) GetById(id string) model.JadwalResponse {
	jadwal, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Jadwal not found",
		})
	}

	response := model.JadwalResponse{
		Id:        jadwal.Id.String(),
		IdPegawai: jadwal.IdPegawai.String(),
		IdHari:    jadwal.IdHari,
		IdShift:   jadwal.IdShift,
		JamMasuk:  jadwal.JamMasuk,
		JamPulang: jadwal.JamPulang,
	}

	return response
}

func (u *JadwalUseCase) Update(request *model.UpdateJadwalRequest, id, user string) model.JadwalResponse {
	jadwal, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Jadwal not found",
		})
	}

	jadwal.Id = helper.MustParse(id)
	jadwal.IdPegawai = helper.MustParse(request.IdPegawai)
	jadwal.IdHari = request.IdHari
	jadwal.IdShift = request.IdShift
	jadwal.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&jadwal); err != nil {
		exception.PanicIfError(err, "Failed to update jadwal")
	}

	response := model.JadwalResponse{
		Id:        jadwal.Id.String(),
		IdPegawai: jadwal.IdPegawai.String(),
		IdHari:    jadwal.IdHari,
		IdShift:   jadwal.IdShift,
		JamMasuk:  jadwal.JamMasuk,
		JamPulang: jadwal.JamPulang,
	}

	return response
}
