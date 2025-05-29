package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
)

type pemeriksaanRanapRepositoryImpl struct {
	DB *sqlx.DB
}

func NewPemeriksaanRanapRepository(db *sqlx.DB) repository.PemeriksaanRanapRepository {
	return &pemeriksaanRanapRepositoryImpl{DB: db}
}

func (r *pemeriksaanRanapRepositoryImpl) Insert(p *entity.PemeriksaanRanap) error {
	query := `
		INSERT INTO pemeriksaan_ranap (
			no_rawat, tgl_perawatan, jam_rawat, suhu_tubuh, tensi, nadi, 
			respirasi, tinggi, berat, spo2, gcs, kesadaran, keluhan, 
			pemeriksaan, alergi, penilaian, rtl, instruksi, evaluasi, nip
		) VALUES (
			:no_rawat, :tgl_perawatan, :jam_rawat, :suhu_tubuh, :tensi, :nadi, 
			:respirasi, :tinggi, :berat, :spo2, :gcs, :kesadaran, :keluhan, 
			:pemeriksaan, :alergi, :penilaian, :rtl, :instruksi, :evaluasi, :nip
		)
	`
	_, err := r.DB.NamedExec(query, p)
	return err
}

func (r *pemeriksaanRanapRepositoryImpl) FindAll() ([]entity.PemeriksaanRanap, error) {
	var list []entity.PemeriksaanRanap
	query := `SELECT * FROM pemeriksaan_ranap ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *pemeriksaanRanapRepositoryImpl) FindByNomorRawat(nomorRawat string) (entity.PemeriksaanRanap, error) {
	var data entity.PemeriksaanRanap
	query := `SELECT * FROM pemeriksaan_ranap WHERE no_rawat = $1 ORDER BY tgl_perawatan DESC, jam_rawat DESC LIMIT 1`
	err := r.DB.Get(&data, query, nomorRawat)
	return data, err
}

func (r *pemeriksaanRanapRepositoryImpl) FindByTanggal(tanggal string) ([]entity.PemeriksaanRanap, error) {
	var list []entity.PemeriksaanRanap
	query := `SELECT * FROM pemeriksaan_ranap WHERE tgl_perawatan = $1 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, tanggal)
	return list, err
}

func (r *pemeriksaanRanapRepositoryImpl) Update(p *entity.PemeriksaanRanap) error {
	query := `
		UPDATE pemeriksaan_ranap SET 
			suhu_tubuh = :suhu_tubuh, tensi = :tensi, nadi = :nadi, nafas = :respirasi,
			tinggi = :tinggi, berat = :berat, spo2 = :spo2, gcs = :gcs, kesadaran = :kesadaran,
			keluhan = :keluhan, pemeriksaan = :pemeriksaan, alergi = :alergi, penilaian = :penilaian,
			rtl = :rtl, instruksi = :instruksi, evaluasi = :evaluasi, nip = :nip
		WHERE no_rawat = :nomor_rawat AND tgl_perawatan = :tanggal AND jam_rawat = :jam
	`
	_, err := r.DB.NamedExec(query, p)
	return err
}

func (r *pemeriksaanRanapRepositoryImpl) Delete(nomorRawat string) error {
	query := `DELETE FROM pemeriksaan_ranap WHERE no_rawat = $1`
	_, err := r.DB.Exec(query, nomorRawat)
	return err
}

func (r *pemeriksaanRanapRepositoryImpl) GetNamaDokter(kode string) (string, error) {
	var nama string
	query := "SELECT nama_dokter FROM dokter WHERE kode_dokter = $1"
	err := r.DB.Get(&nama, query, kode)
	return nama, err
}

func (r *pemeriksaanRanapRepositoryImpl) CheckDokterExists(kodeDokter string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM dokter WHERE kode_dokter = $1)`
	err := r.DB.QueryRow(query, kodeDokter).Scan(&exists)
	return exists, err
}
