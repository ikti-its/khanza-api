package repository

import "github.com/ikti-its/khanza-api/internal/modules/ref/internal/entity"

type RefRepository interface {
	FindRole() ([]entity.Role, error)
	FindJabatan() ([]entity.Jabatan, error)
	FindDepartemen() ([]entity.Departemen, error)
	FindStatusAktif() ([]entity.StatusAktif, error)
	FindShift() ([]entity.Shift, error)
	FindAlasanCuti() ([]entity.AlasanCuti, error)
}
