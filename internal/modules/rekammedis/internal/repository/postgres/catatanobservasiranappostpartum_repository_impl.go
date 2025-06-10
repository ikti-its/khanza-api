package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
)

type catatanObservasiRanapPostpartumRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCatatanObservasiRanapPostpartumRepository(db *sqlx.DB) repository.CatatanObservasiRanapPostpartumRepository {
	return &catatanObservasiRanapPostpartumRepositoryImpl{DB: db}
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) Insert(data *entity.CatatanObservasiRanapPostpartum) error {
	query := `
		INSERT INTO catatan_observasi_ranap_postpartum (
			no_rawat, tgl_perawatan, jam_rawat, gcs, td, hr, rr, suhu,
			spo2, tfu, kontraksi, perdarahan, keterangan, nip
		) VALUES (
			:no_rawat, :tgl_perawatan, :jam_rawat, :gcs, :td, :hr, :rr, :suhu,
			:spo2, :tfu, :kontraksi, :perdarahan, :keterangan, :nip
		)
	`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) FindAll() ([]entity.CatatanObservasiRanapPostpartum, error) {
	var list []entity.CatatanObservasiRanapPostpartum
	query := `SELECT * FROM catatan_observasi_ranap_postpartum ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanapPostpartum, error) {
	var list []entity.CatatanObservasiRanapPostpartum
	query := `SELECT * FROM catatan_observasi_ranap_postpartum WHERE no_rawat = $1 ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanapPostpartum, error) {
	var list []entity.CatatanObservasiRanapPostpartum
	query := `SELECT * FROM catatan_observasi_ranap_postpartum WHERE tgl_perawatan = $1 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, tanggal)
	return list, err
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanapPostpartum, error) {
	var list []entity.CatatanObservasiRanapPostpartum
	query := `SELECT * FROM catatan_observasi_ranap_postpartum WHERE no_rawat = $1 AND tgl_perawatan = $2 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat, tanggal)
	return list, err
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) Update(data *entity.CatatanObservasiRanapPostpartum) error {
	query := `
		UPDATE catatan_observasi_ranap_postpartum SET
			gcs = :gcs, td = :td, hr = :hr, rr = :rr, suhu = :suhu,
			spo2 = :spo2, tfu = :tfu, kontraksi = :kontraksi,
			perdarahan = :perdarahan, keterangan = :keterangan, nip = :nip
		WHERE no_rawat = :no_rawat AND tgl_perawatan = :tgl_perawatan AND jam_rawat = :jam_rawat
	`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *catatanObservasiRanapPostpartumRepositoryImpl) Delete(noRawat string, tglPerawatan string, jamRawat string) error {
	query := `DELETE FROM catatan_observasi_ranap_postpartum WHERE no_rawat = $1 AND tgl_perawatan = $2 AND jam_rawat = $3`
	_, err := r.DB.Exec(query, noRawat, tglPerawatan, jamRawat)
	return err
}
