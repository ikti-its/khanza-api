package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
)

type catatanObservasiRanapKebidananRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCatatanObservasiRanapKebidananRepository(db *sqlx.DB) repository.CatatanObservasiRanapKebidananRepository {
	return &catatanObservasiRanapKebidananRepositoryImpl{DB: db}
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) Insert(data *entity.CatatanObservasiRanapKebidanan) error {
	query := `
		INSERT INTO catatan_observasi_ranap_kebidanan (
			no_rawat, tgl_perawatan, jam_rawat, gcs, td, hr, rr, suhu, spo2,
			kontraksi, bjj, ppv, vt, nip
		) VALUES (
			:no_rawat, :tgl_perawatan, :jam_rawat, :gcs, :td, :hr, :rr, :suhu, :spo2,
			:kontraksi, :bjj, :ppv, :vt, :nip
		)`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindAll() ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan WHERE no_rawat = $1 ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan WHERE tgl_perawatan = $1 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, tanggal)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan WHERE no_rawat = $1 AND tgl_perawatan = $2 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat, tanggal)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) Update(data *entity.CatatanObservasiRanapKebidanan) error {
	query := `
		UPDATE catatan_observasi_ranap_kebidanan SET
			gcs = :gcs, td = :td, hr = :hr, rr = :rr, suhu = :suhu, spo2 = :spo2,
			kontraksi = :kontraksi, bjj = :bjj, ppv = :ppv, vt = :vt, nip = :nip
		WHERE no_rawat = :no_rawat AND tgl_perawatan = :tgl_perawatan AND jam_rawat = :jam_rawat
	`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) Delete(noRawat string, tglPerawatan string, jamRawat string) error {
	query := `DELETE FROM catatan_observasi_ranap_kebidanan WHERE no_rawat = $1 AND tgl_perawatan = $2 AND jam_rawat = $3`
	_, err := r.DB.Exec(query, noRawat, tglPerawatan, jamRawat)
	return err
}
