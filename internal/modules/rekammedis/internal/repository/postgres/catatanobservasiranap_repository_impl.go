package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
)

type catatanObservasiRanapRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCatatanObservasiRanapRepository(db *sqlx.DB) repository.CatatanObservasiRanapRepository {
	return &catatanObservasiRanapRepositoryImpl{DB: db}
}

func (r *catatanObservasiRanapRepositoryImpl) Insert(data *entity.CatatanObservasiRanap) error {
	query := `
		INSERT INTO catatan_observasi_ranap (
			no_rawat, tgl_perawatan, jam_rawat, gcs, td, hr, rr, suhu, spo2, nip
		) VALUES (
			:no_rawat, :tgl_perawatan, :jam_rawat, :gcs, :td, :hr, :rr, :suhu, :spo2, :nip
		)`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *catatanObservasiRanapRepositoryImpl) FindAll() ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap WHERE no_rawat = $1 ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap WHERE tgl_perawatan = $1 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, tanggal)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap WHERE no_rawat = $1 AND tgl_perawatan = $2 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat, tanggal)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) Update(data *entity.CatatanObservasiRanap) error {
	query := `
		UPDATE catatan_observasi_ranap SET
			gcs = :gcs, td = :td, hr = :hr, rr = :rr, suhu = :suhu, spo2 = :spo2, nip = :nip
		WHERE no_rawat = :no_rawat AND tgl_perawatan = :tgl_perawatan AND jam_rawat = :jam_rawat
	`
	_, err := r.DB.NamedExec(query, data)
	return err
}

func (r *catatanObservasiRanapRepositoryImpl) Delete(noRawat string, tglPerawatan string, jamRawat string) error {
	query := `DELETE FROM catatan_observasi_ranap WHERE no_rawat = $1 AND tgl_perawatan = $2 AND jam_rawat = $3`
	_, err := r.DB.Exec(query, noRawat, tglPerawatan, jamRawat)
	return err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByNoRawatAndTanggal2(noRawat string, tanggal string) (*entity.CatatanObservasiRanap, error) {
	query := `
        SELECT*
		FROM catatan_observasi_ranap
        WHERE no_rawat = $1 AND tgl_perawatan = $2
        LIMIT 1
    `
	fmt.Println("📦 Executing query for no_rawat =", noRawat, "tgl =", tanggal)
	var result entity.CatatanObservasiRanap
	if err := r.DB.Get(&result, query, noRawat, tanggal); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("catatan observasi tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal query observasi: %v", err)
	}

	return &result, nil
}

func (r *catatanObservasiRanapRepositoryImpl) UpdateByNoRawatAndTanggal(noRawat string, tgl string, e *entity.CatatanObservasiRanap) error {
	query := `
		UPDATE catatan_observasi_ranap
		SET 
			jam_rawat = $1,
			gcs = $2,
			td = $3,
			hr = $4,
			rr = $5,
			suhu = $6
		WHERE no_rawat = $7 AND tgl_perawatan = $8
	`

	_, err := r.DB.Exec(
		query,
		e.JamRawat,
		e.GCS,
		e.TD,
		e.HR,
		e.RR,
		e.Suhu,
		noRawat,
		e.TglPerawatan,
	)
	fmt.Println("🔧 Updating catatan_observasi for", noRawat, tgl)
	fmt.Printf("➡️  Data: %+v\n", e)
	if err != nil {
		return fmt.Errorf("gagal update catatan observasi: %v", err)
	}

	return nil
}
