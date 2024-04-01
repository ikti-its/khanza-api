package usecase

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/helper"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran/internal/repository"
)

type KehadiranUseCase struct {
	Repository repository.KehadiranRepository
}

func NewKehadiranUseCase(repository *repository.KehadiranRepository) *KehadiranUseCase {
	return &KehadiranUseCase{
		Repository: *repository,
	}
}

func (u *KehadiranUseCase) Attend(request *model.AttendKehadiranRequest) model.KehadiranResponse {
	kehadiran := entity.Kehadiran{
		IdPegawai:       helper.MustParse(request.IdPegawai),
		IdJadwalPegawai: helper.MustParse(request.IdJadwalPegawai),
		Tanggal:         helper.ParseTime(request.Tanggal, "2006-01-02"),
		JamMasuk:        helper.ParseNow(),
	}

	err := u.Repository.Insert(&kehadiran)
	if err != nil {
		exception.PanicIfError(err, "Failed to attend kehadiran")
	}

	response := model.KehadiranResponse{
		Id:        kehadiran.Id.String(),
		IdPegawai: kehadiran.IdPegawai.String(),
		Tanggal:   kehadiran.Tanggal.Format("2006-01-02"),
		JamMasuk:  helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
	}

	return response
}

func (u *KehadiranUseCase) Leave(request *model.LeaveKehadiranRequest, user string) model.KehadiranResponse {
	kehadiran, err := u.Repository.FindById(helper.MustParse(request.Id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	kehadiran.JamPulang = helper.ParseNow()
	kehadiran.Updater = helper.MustParse(user)

	err = u.Repository.Update(&kehadiran)
	if err != nil {
		exception.PanicIfError(err, "Failed to leave kehadiran")
	}

	response := model.KehadiranResponse{
		Id:         kehadiran.Id.String(),
		IdPegawai:  kehadiran.IdPegawai.String(),
		Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
		JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
		JamPulang:  helper.FormatTime(kehadiran.JamPulang, "15:04:05 +07:00"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (u *KehadiranUseCase) Get() []model.KehadiranResponse {
	kehadiran, err := u.Repository.Find()
	if err != nil {
		exception.PanicIfError(err, "Failed to get kehadiran")
	}

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			Id:         kehadiran.Id.String(),
			IdPegawai:  kehadiran.IdPegawai.String(),
			Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
			JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			JamPulang:  helper.FormatTime(kehadiran.JamPulang, "15:04:05 +07:00"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (u *KehadiranUseCase) GetById(id string) model.KehadiranResponse {
	kehadiran, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	response := model.KehadiranResponse{
		Id:         kehadiran.Id.String(),
		IdPegawai:  kehadiran.IdPegawai.String(),
		Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
		JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
		JamPulang:  helper.FormatTime(kehadiran.JamPulang, "15:04:05 +07:00"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (u *KehadiranUseCase) GetByPegawaiId(id string) []model.KehadiranResponse {
	kehadiran, err := u.Repository.FindByPegawaiId(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			Id:         kehadiran.Id.String(),
			IdPegawai:  kehadiran.IdPegawai.String(),
			Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
			JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			JamPulang:  helper.FormatTime(kehadiran.JamPulang, "15:04:05 +07:00"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (u *KehadiranUseCase) GetByTanggal(tanggal string) []model.KehadiranResponse {
	kehadiran, err := u.Repository.FindByTanggal(tanggal)
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			Id:         kehadiran.Id.String(),
			IdPegawai:  kehadiran.IdPegawai.String(),
			Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
			JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			JamPulang:  helper.FormatTime(kehadiran.JamPulang, "15:04:05 +07:00"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (u *KehadiranUseCase) Update(request *model.UpdateKehadiranRequest, id, user string) model.KehadiranResponse {
	kehadiran, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	kehadiran.Id = helper.MustParse(id)
	kehadiran.IdPegawai = helper.MustParse(request.IdPegawai)
	kehadiran.IdJadwalPegawai = helper.MustParse(request.IdJadwalPegawai)
	kehadiran.Tanggal = helper.ParseTime(request.Tanggal, "2006-01-02")
	kehadiran.JamMasuk = helper.ParseTime(request.JamMasuk, "15:04:05")
	kehadiran.JamPulang = helper.ParseTime(request.JamPulang, "15:04:05")
	kehadiran.Keterangan = request.Keterangan
	kehadiran.Updater = helper.MustParse(user)

	if err := u.Repository.Update(&kehadiran); err != nil {
		exception.PanicIfError(err, "Failed to update kehadiran")
	}

	response := model.KehadiranResponse{
		Id:         kehadiran.Id.String(),
		IdPegawai:  kehadiran.IdPegawai.String(),
		Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
		JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
		JamPulang:  helper.FormatTime(kehadiran.JamPulang, "15:04:05 +07:00"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}
