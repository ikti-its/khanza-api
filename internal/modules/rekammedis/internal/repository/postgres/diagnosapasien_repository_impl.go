package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
)

type diagnosaPasienRepositoryImpl struct {
	DB *sqlx.DB
}

func NewDiagnosaPasienRepository(db *sqlx.DB) repository.DiagnosaPasienRepository {
	return &diagnosaPasienRepositoryImpl{DB: db}
}

func (r *diagnosaPasienRepositoryImpl) Insert(data *entity.DiagnosaPasien) error {
	query := `
		INSERT INTO diagnosa_pasien (
			no_rawat, kd_penyakit, status, prioritas, status_penyakit
		) VALUES (
			:no_rawat, :kd_penyakit, :status, :prioritas, :status_penyakit
		)`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *diagnosaPasienRepositoryImpl) FindAll() ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien ORDER BY no_rawat, prioritas ASC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien WHERE no_rawat = $1 ORDER BY prioritas ASC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) FindByKodePenyakit(kode string) ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien WHERE kd_penyakit = $1`
	err := r.DB.Select(&list, query, kode)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) FindByNoRawatAndStatus(noRawat string, status string) ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien WHERE no_rawat = $1 AND status = $2 ORDER BY prioritas ASC`
	err := r.DB.Select(&list, query, noRawat, status)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) Update(data *entity.DiagnosaPasien) error {
	query := `
		UPDATE diagnosa_pasien SET
			status = :status,
			prioritas = :prioritas,
			status_penyakit = :status_penyakit
		WHERE no_rawat = :no_rawat AND kd_penyakit = :kd_penyakit
	`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *diagnosaPasienRepositoryImpl) Delete(noRawat string, kdPenyakit string) error {
	query := `DELETE FROM diagnosa_pasien WHERE no_rawat = $1 AND kd_penyakit = $2`
	_, err := r.DB.Exec(query, noRawat, kdPenyakit)
	return err
}
