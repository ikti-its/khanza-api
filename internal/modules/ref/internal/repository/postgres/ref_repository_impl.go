package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/ref/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/ref/internal/repository"
	"github.com/jmoiron/sqlx"
)

type refRepositoryImpl struct {
	DB *sqlx.DB
}

func (r refRepositoryImpl) FindRole() ([]entity.Role, error) {
	query := `
		SELECT id, nama
		FROM ref.role
	`

	var records []entity.Role
	err := r.DB.Select(&records, query)

	return records, err
}

func (r refRepositoryImpl) FindJabatan() ([]entity.Jabatan, error) {
	query := `
		SELECT id, nama
		FROM ref.jabatan
	`

	var records []entity.Jabatan
	err := r.DB.Select(&records, query)

	return records, err
}

func (r refRepositoryImpl) FindDepartemen() ([]entity.Departemen, error) {
	query := `
		SELECT id, nama
		FROM ref.departemen
	`

	var records []entity.Departemen
	err := r.DB.Select(&records, query)

	return records, err
}

func (r refRepositoryImpl) FindStatusAktif() ([]entity.StatusAktif, error) {
	query := `
		SELECT id, nama
		FROM ref.status_aktif_pegawai
	`

	var records []entity.StatusAktif
	err := r.DB.Select(&records, query)

	return records, err
}

func (r refRepositoryImpl) FindShift() ([]entity.Shift, error) {
	query := `
		SELECT id, nama
		FROM ref.shift
	`

	var records []entity.Shift
	err := r.DB.Select(&records, query)

	return records, err
}

func (r refRepositoryImpl) FindAlasanCuti() ([]entity.AlasanCuti, error) {
	query := `
		SELECT id, nama
		FROM ref.alasan_cuti
	`

	var records []entity.AlasanCuti
	err := r.DB.Select(&records, query)

	return records, err
}

func NewRefRepository(db *sqlx.DB) repository.RefRepository {
	return &refRepositoryImpl{db}
}
