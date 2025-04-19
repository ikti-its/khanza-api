package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/repository"
	"github.com/jmoiron/sqlx"
)

type pemberianObatRepositoryImpl struct {
	DB *sqlx.DB
}

func NewPemberianObatRepository(db *sqlx.DB) repository.PemberianObatRepository {
	return &pemberianObatRepositoryImpl{DB: db}
}

func (r *pemberianObatRepositoryImpl) Insert(p *entity.PemberianObat) error {
	query := `
		INSERT INTO pemberian_obat (
			tanggal_beri, jam_beri, nomor_rawat, nama_pasien, kode_obat, 
			nama_obat, embalase, tuslah, jumlah, biaya_obat, total, 
			gudang, no_batch, no_faktur
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10, $11,
			$12, $13, $14
		)
	`
	_, err := r.DB.Exec(query,
		p.TanggalBeri, p.JamBeri, p.NomorRawat, p.NamaPasien, p.KodeObat,
		p.NamaObat, p.Embalase, p.Tuslah, p.Jumlah, p.BiayaObat, p.Total,
		p.Gudang, p.NoBatch, p.NoFaktur,
	)
	return err
}

func (r *pemberianObatRepositoryImpl) FindAll() ([]entity.PemberianObat, error) {
	query := `SELECT * FROM pemberian_obat ORDER BY tanggal_beri DESC, jam_beri DESC`
	var list []entity.PemberianObat
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *pemberianObatRepositoryImpl) FindByNomorRawat(nomorRawat string) ([]entity.PemberianObat, error) {
	query := `SELECT * FROM pemberian_obat WHERE nomor_rawat = $1 ORDER BY tanggal_beri DESC, jam_beri DESC`
	var list []entity.PemberianObat
	err := r.DB.Select(&list, query, nomorRawat)
	return list, err
}

func (r *pemberianObatRepositoryImpl) Update(p *entity.PemberianObat) error {
	query := `
		UPDATE pemberian_obat SET 
			nama_pasien = $4, kode_obat = $5, nama_obat = $6, embalase = $7,
			tuslah = $8, jumlah = $9, biaya_obat = $10, total = $11,
			gudang = $12, no_batch = $13, no_faktur = $14
		WHERE nomor_rawat = $3 AND tanggal_beri = $1 AND jam_beri = $2
	`
	_, err := r.DB.Exec(query,
		p.TanggalBeri, p.JamBeri, p.NomorRawat, p.NamaPasien, p.KodeObat,
		p.NamaObat, p.Embalase, p.Tuslah, p.Jumlah, p.BiayaObat, p.Total,
		p.Gudang, p.NoBatch, p.NoFaktur,
	)
	return err
}

func (r *pemberianObatRepositoryImpl) Delete(nomorRawat, jamBeri string) error {
	query := `DELETE FROM pemberian_obat WHERE nomor_rawat = $1 AND jam_beri = $2`
	_, err := r.DB.Exec(query, nomorRawat, jamBeri)
	return err
}
