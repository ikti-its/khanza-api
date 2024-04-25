package usecase

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/repository"
	"log"
)

type HomeUseCase struct {
	Repository repository.HomeRepository
}

func NewHomeUseCase(repository *repository.HomeRepository) *HomeUseCase {
	return &HomeUseCase{
		Repository: *repository,
	}
}

func (u *HomeUseCase) GetHomePegawai(id string, hari int, tanggal string) model.HomeResponse {
	if hari == 0 || tanggal == "" || helper.ParseTime(tanggal, "2006-01-02").IsZero() {
		panic(&exception.BadRequestError{
			Message: "Invalid query format",
		})
	}

	home, err := u.Repository.HomePegawai(helper.MustParse(id), hari)
	if err != nil {
		log.Printf(err.Error())
		panic(&exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	status := "Hadir"
	if presensi, err := u.Repository.ObserveKehadiran(home.Pegawai, home.Jadwal, tanggal); presensi == uuid.Nil && err != nil {
		status = "Belum Hadir"
	}

	response := model.HomeResponse{
		Akun:      home.Akun.String(),
		Pegawai:   home.Pegawai.String(),
		Nama:      home.Nama,
		NIP:       home.NIP,
		Profil:    home.Profil,
		Alamat:    home.Alamat,
		AlamatLat: home.AlamatLat,
		AlamatLon: home.AlamatLon,
		Foto:      home.Foto,
		Shift:     home.Shift,
		JamMasuk:  helper.FormatTime(home.JamMasuk, "15:04"),
		JamPulang: helper.FormatTime(home.JamPulang, "15:04"),
		Status:    status,
	}

	return response
}
