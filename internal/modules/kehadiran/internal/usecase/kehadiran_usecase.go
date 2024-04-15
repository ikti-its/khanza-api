package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/repository"
)

type KehadiranUseCase struct {
	Repository repository.KehadiranRepository
}

func NewKehadiranUseCase(repository *repository.KehadiranRepository) *KehadiranUseCase {
	return &KehadiranUseCase{
		Repository: *repository,
	}
}

func (u *KehadiranUseCase) Attend(request *model.AttendKehadiranRequest, updater string) model.KehadiranResponse {
	kehadiran := entity.Kehadiran{
		Id:              helper.MustNew(),
		IdPegawai:       helper.MustParse(request.IdPegawai),
		IdJadwalPegawai: helper.MustParse(request.IdJadwalPegawai),
		Tanggal:         helper.ParseTime(request.Tanggal, "2006-01-02"),
		JamMasuk:        helper.ParseNow(),
		Updater:         helper.MustParse(updater),
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

func (u *KehadiranUseCase) Leave(request *model.LeaveKehadiranRequest, updater string) model.KehadiranResponse {
	kehadiran, err := u.Repository.FindById(helper.MustParse(request.Id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	kehadiran.JamPulang.Time = helper.ParseNow()
	kehadiran.JamPulang.Valid = true
	kehadiran.Updater = helper.MustParse(updater)

	err = u.Repository.Update(&kehadiran)
	if err != nil {
		exception.PanicIfError(err, "Failed to leave kehadiran")
	}

	response := model.KehadiranResponse{
		Id:         kehadiran.Id.String(),
		IdPegawai:  kehadiran.IdPegawai.String(),
		Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
		JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
		JamPulang:  helper.FormatTime(kehadiran.JamPulang.Time, "15:04:05 +07:00"),
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
		if kehadiran.JamPulang.Valid {
			response[i] = model.KehadiranResponse{
				Id:         kehadiran.Id.String(),
				IdPegawai:  kehadiran.IdPegawai.String(),
				Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
				JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
				JamPulang:  helper.FormatTime(kehadiran.JamPulang.Time, "15:04:05 +07:00"),
				Keterangan: kehadiran.Keterangan,
			}
		} else {
			response[i] = model.KehadiranResponse{
				Id:        kehadiran.Id.String(),
				IdPegawai: kehadiran.IdPegawai.String(),
				Tanggal:   kehadiran.Tanggal.Format("2006-01-02"),
				JamMasuk:  helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			}
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

	if kehadiran.JamPulang.Valid {
		response := model.KehadiranResponse{
			Id:         kehadiran.Id.String(),
			IdPegawai:  kehadiran.IdPegawai.String(),
			Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
			JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			JamPulang:  helper.FormatTime(kehadiran.JamPulang.Time, "15:04:05 +07:00"),
			Keterangan: kehadiran.Keterangan,
		}

		return response
	} else {
		response := model.KehadiranResponse{
			Id:        kehadiran.Id.String(),
			IdPegawai: kehadiran.IdPegawai.String(),
			Tanggal:   kehadiran.Tanggal.Format("2006-01-02"),
			JamMasuk:  helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
		}

		return response
	}
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
		if kehadiran.JamPulang.Valid {
			response[i] = model.KehadiranResponse{
				Id:         kehadiran.Id.String(),
				IdPegawai:  kehadiran.IdPegawai.String(),
				Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
				JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
				JamPulang:  helper.FormatTime(kehadiran.JamPulang.Time, "15:04:05 +07:00"),
				Keterangan: kehadiran.Keterangan,
			}
		} else {
			response[i] = model.KehadiranResponse{
				Id:        kehadiran.Id.String(),
				IdPegawai: kehadiran.IdPegawai.String(),
				Tanggal:   kehadiran.Tanggal.Format("2006-01-02"),
				JamMasuk:  helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			}
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
		if kehadiran.JamPulang.Valid {
			response[i] = model.KehadiranResponse{
				Id:         kehadiran.Id.String(),
				IdPegawai:  kehadiran.IdPegawai.String(),
				Tanggal:    kehadiran.Tanggal.Format("2006-01-02"),
				JamMasuk:   helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
				JamPulang:  helper.FormatTime(kehadiran.JamPulang.Time, "15:04:05 +07:00"),
				Keterangan: kehadiran.Keterangan,
			}
		} else {
			response[i] = model.KehadiranResponse{
				Id:        kehadiran.Id.String(),
				IdPegawai: kehadiran.IdPegawai.String(),
				Tanggal:   kehadiran.Tanggal.Format("2006-01-02"),
				JamMasuk:  helper.FormatTime(kehadiran.JamMasuk, "15:04:05 +07:00"),
			}
		}
	}

	return response
}
