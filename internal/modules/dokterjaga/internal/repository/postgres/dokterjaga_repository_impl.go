package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/dokterjaga/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DokterJagaRepository interface {
	Insert(dokter *entity.DokterJaga) error
	FindAll() ([]entity.DokterJaga, error)
	FindByKodeDokter(kodeDokter string) ([]entity.DokterJaga, error)
	Update(dokter *entity.DokterJaga) error
	Delete(kodeDokter string, hariKerja string) error
	FindByStatus(status string) ([]entity.DokterJaga, error)
	UpdateStatus(kodeDokter string, hariKerja string, status string) error
}

type dokterJagaRepositoryImpl struct {
	DB *sqlx.DB
}

func NewDokterJagaRepository(db *sqlx.DB) DokterJagaRepository {
	return &dokterJagaRepositoryImpl{DB: db}
}

func (r *dokterJagaRepositoryImpl) Insert(d *entity.DokterJaga) error {
	query := `
		INSERT INTO dokter_jaga (
			kode_dokter, nama_dokter, hari_kerja,
			jam_mulai, jam_selesai, poliklinik, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(query,
		d.KodeDokter, d.NamaDokter, d.HariKerja,
		d.JamMulai, d.JamSelesai, d.Poliklinik, d.Status,
	)
	return err
}

func (r *dokterJagaRepositoryImpl) FindAll() ([]entity.DokterJaga, error) {
	query := `SELECT * FROM dokter_jaga ORDER BY hari_kerja DESC`
	var records []entity.DokterJaga
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *dokterJagaRepositoryImpl) FindByKodeDokter(kode string) ([]entity.DokterJaga, error) {
	query := `SELECT * FROM dokter_jaga WHERE kode_dokter = $1 ORDER BY hari_kerja DESC`
	var records []entity.DokterJaga
	err := r.DB.Select(&records, query, kode)
	return records, err
}

func (r *dokterJagaRepositoryImpl) Update(d *entity.DokterJaga) error {
	query := `
		UPDATE dokter_jaga SET 
			nama_dokter = $2,
			jam_mulai = $3,
			jam_selesai = $4,
			poliklinik = $5,
			status = $6
		WHERE kode_dokter = $1 AND hari_kerja = $7
	`
	_, err := r.DB.Exec(query,
		d.KodeDokter,
		d.NamaDokter,
		d.JamMulai,
		d.JamSelesai,
		d.Poliklinik,
		d.Status,
		d.HariKerja,
	)
	return err
}

func (r *dokterJagaRepositoryImpl) Delete(kodeDokter string, hariKerja string) error {
	query := `DELETE FROM dokter_jaga WHERE kode_dokter = $1 AND hari_kerja = $2`
	_, err := r.DB.Exec(query, kodeDokter, hariKerja)
	return err
}

func (r *dokterJagaRepositoryImpl) FindByStatus(status string) ([]entity.DokterJaga, error) {
	query := `SELECT * FROM dokter_jaga WHERE status = $1 ORDER BY hari_kerja DESC`
	var records []entity.DokterJaga
	err := r.DB.Select(&records, query, status)
	return records, err
}

func (r *dokterJagaRepositoryImpl) UpdateStatus(kodeDokter string, hariKerja string, status string) error {
	query := `UPDATE dokter_jaga SET status = $1 WHERE kode_dokter = $2 AND hari_kerja = $3`
	_, err := r.DB.Exec(query, status, kodeDokter, hariKerja)
	return err
}
