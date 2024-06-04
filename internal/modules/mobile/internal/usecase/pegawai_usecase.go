package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/app/helper"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/repository"
)

type PegawaiUseCase struct {
	Repository repository.PegawaiRepository
}

func NewPegawaiUseCase(repository *repository.PegawaiRepository) *PegawaiUseCase {
	return &PegawaiUseCase{
		Repository: *repository,
	}
}

func (u *PegawaiUseCase) Get() []model.PegawaiResponse {
	pegawai, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all pegawai")

	response := make([]model.PegawaiResponse, len(pegawai))
	for i, pegawai := range pegawai {
		response[i] = model.PegawaiResponse{
			Pegawai:      pegawai.Pegawai.String(),
			Akun:         pegawai.Akun.String(),
			NIP:          pegawai.NIP,
			NIK:          pegawai.NIK,
			Nama:         pegawai.Nama,
			JenisKelamin: pegawai.JenisKelamin,
			TempatLahir:  pegawai.TempatLahir,
			TanggalLahir: helper.FormatTime(pegawai.TanggalLahir, "2006-01-02"),
			Agama:        pegawai.Agama,
			Pendidikan:   pegawai.Pendidikan,
			Jabatan:      pegawai.Jabatan,
			Departemen:   pegawai.Departemen,
			Status:       pegawai.Status,
			JenisPegawai: pegawai.JenisPegawai,
			Telepon:      pegawai.Telepon,
			TanggalMasuk: helper.FormatTime(pegawai.TanggalMasuk, "2006-01-02"),
		}
	}

	return response
}

func (u *PegawaiUseCase) GetPage(page, size int) model.PegawaiPageResponse {
	pegawai, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged pegawai")

	response := make([]model.PegawaiResponse, len(pegawai))
	for i, pegawai := range pegawai {
		response[i] = model.PegawaiResponse{
			Pegawai:      pegawai.Pegawai.String(),
			Akun:         pegawai.Akun.String(),
			NIP:          pegawai.NIP,
			NIK:          pegawai.NIK,
			Nama:         pegawai.Nama,
			JenisKelamin: pegawai.JenisKelamin,
			TempatLahir:  pegawai.TempatLahir,
			TanggalLahir: helper.FormatTime(pegawai.TanggalLahir, "2006-01-02"),
			Agama:        pegawai.Agama,
			Pendidikan:   pegawai.Pendidikan,
			Jabatan:      pegawai.Jabatan,
			Departemen:   pegawai.Departemen,
			Status:       pegawai.Status,
			JenisPegawai: pegawai.JenisPegawai,
			Telepon:      pegawai.Telepon,
			TanggalMasuk: helper.FormatTime(pegawai.TanggalMasuk, "2006-01-02"),
		}
	}

	pagedResponse := model.PegawaiPageResponse{
		Page:    page,
		Size:    size,
		Total:   total,
		Pegawai: response,
	}

	return pagedResponse
}

func (u *PegawaiUseCase) GetById(id string) model.PegawaiResponse {
	pegawai, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Pegawai not found",
		})
	}

	response := model.PegawaiResponse{
		Pegawai:      pegawai.Pegawai.String(),
		Akun:         pegawai.Akun.String(),
		NIP:          pegawai.NIP,
		NIK:          pegawai.NIK,
		Nama:         pegawai.Nama,
		JenisKelamin: pegawai.JenisKelamin,
		TempatLahir:  pegawai.TempatLahir,
		TanggalLahir: helper.FormatTime(pegawai.TanggalLahir, "2006-01-02"),
		Agama:        pegawai.Agama,
		Pendidikan:   pegawai.Pendidikan,
		Jabatan:      pegawai.Jabatan,
		Departemen:   pegawai.Departemen,
		Status:       pegawai.Status,
		JenisPegawai: pegawai.JenisPegawai,
		Telepon:      pegawai.Telepon,
		TanggalMasuk: helper.FormatTime(pegawai.TanggalMasuk, "2006-01-02"),
	}

	return response
}
