package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/modules/ref/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/ref/internal/repository"
)

type RefUseCase struct {
	Repository repository.RefRepository
}

func NewRefUseCase(repository *repository.RefRepository) *RefUseCase {
	return &RefUseCase{
		Repository: *repository,
	}
}

func (u *RefUseCase) GetRole() []model.RoleResponse {
	role, err := u.Repository.FindRole()
	exception.PanicIfError(err, "Failed to get all role")

	response := make([]model.RoleResponse, len(role))
	for i, role := range role {
		response[i] = model.RoleResponse{
			Id:   role.Id,
			Nama: role.Nama,
		}
	}

	return response
}

func (u *RefUseCase) GetJabatan() []model.JabatanResponse {
	jabatan, err := u.Repository.FindJabatan()
	exception.PanicIfError(err, "Failed to get all jabatan")

	response := make([]model.JabatanResponse, len(jabatan))
	for i, jabatan := range jabatan {
		response[i] = model.JabatanResponse{
			Id:   jabatan.Id,
			Nama: jabatan.Nama,
		}
	}

	return response
}

func (u *RefUseCase) GetDepartemen() []model.DepartemenResponse {
	departemen, err := u.Repository.FindDepartemen()
	exception.PanicIfError(err, "Failed to get all departemen")

	response := make([]model.DepartemenResponse, len(departemen))
	for i, departemen := range departemen {
		response[i] = model.DepartemenResponse{
			Id:   departemen.Id,
			Nama: departemen.Nama,
		}
	}

	return response
}

func (u *RefUseCase) GetStatusAktif() []model.StatusAktifResponse {
	status, err := u.Repository.FindStatusAktif()
	exception.PanicIfError(err, "Failed to get all status")

	response := make([]model.StatusAktifResponse, len(status))
	for i, status := range status {
		response[i] = model.StatusAktifResponse{
			Id:   status.Id,
			Nama: status.Nama,
		}
	}

	return response
}

func (u *RefUseCase) GetShift() []model.ShiftResponse {
	shift, err := u.Repository.FindShift()
	exception.PanicIfError(err, "Failed to get all shift")

	response := make([]model.ShiftResponse, len(shift))
	for i, shift := range shift {
		response[i] = model.ShiftResponse{
			Id:   shift.Id,
			Nama: shift.Nama,
		}
	}

	return response
}

func (u *RefUseCase) GetAlasanCuti() []model.AlasanCutiResponse {
	alasan, err := u.Repository.FindAlasanCuti()
	exception.PanicIfError(err, "Failed to get all alasan")

	response := make([]model.AlasanCutiResponse, len(alasan))
	for i, alasan := range alasan {
		response[i] = model.AlasanCutiResponse{
			Id:   alasan.Id,
			Nama: alasan.Nama,
		}
	}

	return response
}
