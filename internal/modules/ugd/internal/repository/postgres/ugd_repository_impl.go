package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/ugd/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/ugd/internal/repository"
	"github.com/jmoiron/sqlx"
)

type UGDRepository interface {
	Insert(ugd *entity.UGD) error
	Find() ([]entity.UGD, error)
	FindAll() ([]entity.UGD, error)
	FindByNomorReg(nomorReg string) (entity.UGD, error)
	FindByNomorRM(nomorRM string) (entity.UGD, error)
	FindByTanggal(tanggal string) ([]entity.UGD, error)
	Update(ugd *entity.UGD) error
	Delete(nomorReg string) error
	CheckDokterExists(kodeDokter string) (bool, error)
}

type ugdRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUGDRepository(db *sqlx.DB) repository.UGDRepository {
	return &ugdRepositoryImpl{DB: db}
}

func (r *ugdRepositoryImpl) Insert(ugd *entity.UGD) error {
	query := `
		INSERT INTO ugd (
			nomor_reg, nomor_rawat, tanggal, jam, kode_dokter, dokter_dituju, nomor_rm,
			nama_pasien, jenis_kelamin, umur, poliklinik, jenis_bayar, penanggung_jawab,
			alamat_pj, hubungan_pj, biaya_registrasi, status, status_rawat, status_bayar
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19
		)
	`
	_, err := r.DB.Exec(query,
		ugd.NomorReg, ugd.NomorRawat, ugd.Tanggal, ugd.Jam, ugd.KodeDokter, ugd.DokterDituju,
		ugd.NomorRM, ugd.NamaPasien, ugd.JenisKelamin, ugd.Umur, ugd.Poliklinik,
		ugd.JenisBayar, ugd.PenanggungJawab, ugd.AlamatPJ, ugd.HubunganPJ, ugd.BiayaRegistrasi,
		ugd.Status, ugd.StatusRawat, ugd.StatusBayar,
	)
	return err
}

func (r *ugdRepositoryImpl) Update(ugd *entity.UGD) error {
	query := `
		UPDATE ugd SET 
			nomor_rawat = $2, tanggal = $3, jam = $4, kode_dokter = $5, dokter_dituju = $6,
			nomor_rm = $7, nama_pasien = $8, jenis_kelamin = $9, umur = $10, poliklinik = $11,
			jenis_bayar = $12, penanggung_jawab = $13, alamat_pj = $14, hubungan_pj = $15,
			biaya_registrasi = $16, status = $17, status_rawat = $18, status_bayar = $19
		WHERE nomor_reg = $1
	`
	_, err := r.DB.Exec(query,
		ugd.NomorReg, ugd.NomorRawat, ugd.Tanggal, ugd.Jam, ugd.KodeDokter, ugd.DokterDituju,
		ugd.NomorRM, ugd.NamaPasien, ugd.JenisKelamin, ugd.Umur, ugd.Poliklinik,
		ugd.JenisBayar, ugd.PenanggungJawab, ugd.AlamatPJ, ugd.HubunganPJ, ugd.BiayaRegistrasi,
		ugd.Status, ugd.StatusRawat, ugd.StatusBayar,
	)
	return err
}

func (r *ugdRepositoryImpl) Find() ([]entity.UGD, error) {
	query := `SELECT * FROM ugd ORDER BY tanggal DESC`
	var records []entity.UGD
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *ugdRepositoryImpl) FindAll() ([]entity.UGD, error) {
	return r.Find()
}

func (r *ugdRepositoryImpl) FindByNomorReg(nomorReg string) (entity.UGD, error) {
	query := `SELECT * FROM ugd WHERE nomor_reg = $1`
	var record entity.UGD
	err := r.DB.Get(&record, query, nomorReg)
	return record, err
}

func (r *ugdRepositoryImpl) FindByNomorRM(nomorRM string) (entity.UGD, error) {
	query := `SELECT * FROM ugd WHERE nomor_rm = $1`
	var record entity.UGD
	err := r.DB.Get(&record, query, nomorRM)
	return record, err
}

func (r *ugdRepositoryImpl) FindByTanggal(tanggal string) ([]entity.UGD, error) {
	query := `SELECT * FROM ugd WHERE tanggal = $1`
	var records []entity.UGD
	err := r.DB.Select(&records, query, tanggal)
	return records, err
}

func (r *ugdRepositoryImpl) Delete(nomorReg string) error {
	query := `DELETE FROM ugd WHERE nomor_reg = $1`
	_, err := r.DB.Exec(query, nomorReg)
	return err
}

func (r *ugdRepositoryImpl) CheckDokterExists(kodeDokter string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM dokter WHERE kode_dokter = $1)`
	err := r.DB.QueryRow(query, kodeDokter).Scan(&exists)
	return exists, err
}
