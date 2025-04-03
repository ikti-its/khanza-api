package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/entity"
	"github.com/jmoiron/sqlx"
)

type RujukanMasukRepository interface {
	Insert(rujukan *entity.RujukanMasuk) error
	FindAll() ([]entity.RujukanMasuk, error)
	FindByNomorRawat(nomorRawat string) (entity.RujukanMasuk, error)
	FindByNomorRM(nomorRM string) ([]entity.RujukanMasuk, error)
	FindByTanggalMasuk(tanggal string) ([]entity.RujukanMasuk, error)
	Update(rujukan *entity.RujukanMasuk) error
	Delete(nomorRawat string) error
}

type rujukanMasukRepositoryImpl struct {
	DB *sqlx.DB
}

func NewRujukanMasukRepository(db *sqlx.DB) RujukanMasukRepository {
	return &rujukanMasukRepositoryImpl{DB: db}
}

func (r *rujukanMasukRepositoryImpl) Insert(rujukan *entity.RujukanMasuk) error {
	query := `
		INSERT INTO rujukan_masuk (
			nomor_rujuk, perujuk, alamat_perujuk, nomor_rawat,
			nomor_rm, nama_pasien, alamat, umur,
			tanggal_masuk, tanggal_keluar, diagnosa_awal
		) VALUES (
			$1, $2, $3, $4,
			$5, $6, $7, $8,
			$9, $10, $11
		)
	`
	_, err := r.DB.Exec(query,
		rujukan.NomorRujuk, rujukan.Perujuk, rujukan.AlamatPerujuk, rujukan.NomorRawat,
		rujukan.NomorRM, rujukan.NamaPasien, rujukan.Alamat, rujukan.Umur,
		rujukan.TanggalMasuk, rujukan.TanggalKeluar, rujukan.DiagnosaAwal,
	)
	return err
}

func (r *rujukanMasukRepositoryImpl) FindAll() ([]entity.RujukanMasuk, error) {
	query := `SELECT * FROM rujukan_masuk ORDER BY tanggal_masuk DESC`
	var records []entity.RujukanMasuk
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *rujukanMasukRepositoryImpl) FindByNomorRawat(nomorRawat string) (entity.RujukanMasuk, error) {
	query := `SELECT * FROM rujukan_masuk WHERE nomor_rawat = $1`
	var record entity.RujukanMasuk
	err := r.DB.Get(&record, query, nomorRawat)
	return record, err
}

func (r *rujukanMasukRepositoryImpl) FindByNomorRM(nomorRM string) ([]entity.RujukanMasuk, error) {
	query := `SELECT * FROM rujukan_masuk WHERE nomor_rm = $1 ORDER BY tanggal_masuk DESC`
	var records []entity.RujukanMasuk
	err := r.DB.Select(&records, query, nomorRM)
	return records, err
}

func (r *rujukanMasukRepositoryImpl) FindByTanggalMasuk(tanggal string) ([]entity.RujukanMasuk, error) {
	query := `SELECT * FROM rujukan_masuk WHERE tanggal_masuk = $1 ORDER BY nomor_rawat`
	var records []entity.RujukanMasuk
	err := r.DB.Select(&records, query, tanggal)
	return records, err
}

func (r *rujukanMasukRepositoryImpl) Update(rujukan *entity.RujukanMasuk) error {
	query := `
		UPDATE rujukan_masuk SET 
			nomor_rujuk = $1,
			perujuk = $2,
			alamat_perujuk = $3,
			nomor_rm = $4,
			nama_pasien = $5,
			alamat = $6,
			umur = $7,
			tanggal_masuk = $8,
			tanggal_keluar = $9,
			diagnosa_awal = $10
		WHERE nomor_rawat = $11
	`
	_, err := r.DB.Exec(query,
		rujukan.NomorRujuk, rujukan.Perujuk, rujukan.AlamatPerujuk,
		rujukan.NomorRM, rujukan.NamaPasien, rujukan.Alamat,
		rujukan.Umur, rujukan.TanggalMasuk, rujukan.TanggalKeluar,
		rujukan.DiagnosaAwal, rujukan.NomorRawat,
	)
	return err
}

func (r *rujukanMasukRepositoryImpl) Delete(nomorRawat string) error {
	query := `DELETE FROM rujukan_masuk WHERE nomor_rawat = $1`
	_, err := r.DB.Exec(query, nomorRawat)
	return err
}
