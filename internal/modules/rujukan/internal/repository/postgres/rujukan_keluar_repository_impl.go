package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/entity"
	"github.com/jmoiron/sqlx"
)

type RujukanKeluarRepository interface {
	Insert(rujukan *entity.RujukanKeluar) error
	FindAll() ([]entity.RujukanKeluar, error)
	FindByNomorRawat(nomorRawat string) (entity.RujukanKeluar, error)
	FindByNomorRM(nomorRM string) ([]entity.RujukanKeluar, error)
	FindByTanggalRujuk(tanggal string) ([]entity.RujukanKeluar, error)
	Update(rujukan *entity.RujukanKeluar) error
	Delete(nomorRawat string) error
}

type rujukanKeluarRepositoryImpl struct {
	DB *sqlx.DB
}

func NewRujukanKeluarRepository(db *sqlx.DB) RujukanKeluarRepository {
	return &rujukanKeluarRepositoryImpl{DB: db}
}

func (r *rujukanKeluarRepositoryImpl) Insert(rujukan *entity.RujukanKeluar) error {
	query := `
		INSERT INTO rujukan_keluar (
			nomor_rujuk, nomor_rawat, nomor_rm, nama_pasien,
			tempat_rujuk, tanggal_rujuk, jam_rujuk,
			keterangan_diagnosa, dokter_perujuk, kategori_rujuk,
			pengantaran, keterangan
		) VALUES (
			$1, $2, $3, $4,
			$5, $6, $7,
			$8, $9, $10,
			$11, $12
		)
	`
	_, err := r.DB.Exec(query,
		rujukan.NomorRujuk, rujukan.NomorRawat, rujukan.NomorRM, rujukan.NamaPasien,
		rujukan.TempatRujuk, rujukan.TanggalRujuk, rujukan.JamRujuk,
		rujukan.KeteranganDiagnosa, rujukan.DokterPerujuk, rujukan.KategoriRujuk,
		rujukan.Pengantaran, rujukan.Keterangan,
	)
	return err
}

func (r *rujukanKeluarRepositoryImpl) FindAll() ([]entity.RujukanKeluar, error) {
	query := `SELECT * FROM rujukan_keluar ORDER BY tanggal_rujuk DESC`
	var records []entity.RujukanKeluar
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *rujukanKeluarRepositoryImpl) FindByNomorRawat(nomorRawat string) (entity.RujukanKeluar, error) {
	query := `SELECT * FROM rujukan_keluar WHERE nomor_rawat = $1`
	var record entity.RujukanKeluar
	err := r.DB.Get(&record, query, nomorRawat)
	return record, err
}

func (r *rujukanKeluarRepositoryImpl) FindByNomorRM(nomorRM string) ([]entity.RujukanKeluar, error) {
	query := `SELECT * FROM rujukan_keluar WHERE nomor_rm = $1 ORDER BY tanggal_rujuk DESC`
	var records []entity.RujukanKeluar
	err := r.DB.Select(&records, query, nomorRM)
	return records, err
}

func (r *rujukanKeluarRepositoryImpl) FindByTanggalRujuk(tanggal string) ([]entity.RujukanKeluar, error) {
	query := `SELECT * FROM rujukan_keluar WHERE tanggal_rujuk = $1 ORDER BY nomor_rawat`
	var records []entity.RujukanKeluar
	err := r.DB.Select(&records, query, tanggal)
	return records, err
}

func (r *rujukanKeluarRepositoryImpl) Update(rujukan *entity.RujukanKeluar) error {
	query := `
		UPDATE rujukan_keluar SET 
			nomor_rujuk = $1,
			nomor_rm = $2,
			nama_pasien = $3,
			tempat_rujuk = $4,
			tanggal_rujuk = $5,
			jam_rujuk = $6,
			keterangan_diagnosa = $7,
			dokter_perujuk = $8,
			kategori_rujuk = $9,
			pengantaran = $10,
			keterangan = $11
		WHERE nomor_rawat = $12
	`
	_, err := r.DB.Exec(query,
		rujukan.NomorRujuk, rujukan.NomorRM, rujukan.NamaPasien, rujukan.TempatRujuk,
		rujukan.TanggalRujuk, rujukan.JamRujuk, rujukan.KeteranganDiagnosa,
		rujukan.DokterPerujuk, rujukan.KategoriRujuk, rujukan.Pengantaran,
		rujukan.Keterangan, rujukan.NomorRawat,
	)
	return err
}

func (r *rujukanKeluarRepositoryImpl) Delete(nomorRawat string) error {
	query := `DELETE FROM rujukan_keluar WHERE nomor_rawat = $1`
	_, err := r.DB.Exec(query, nomorRawat)
	return err
}
