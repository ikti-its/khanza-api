package postgres

import (
	"fmt"
	"log"

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
			suhu_tubuh = :suhu_tubuh, tensi = :tensi, nadi = :nadi, respirasi = :respirasi,
			tinggi = :tinggi, berat = :berat, spo2 = :spo2, gcs = :gcs, kesadaran = :kesadaran,
			keluhan = :keluhan, pemeriksaan = :pemeriksaan, alergi = :alergi, penilaian = :penilaian,
			rtl = :rtl, instruksi = :instruksi, evaluasi = :evaluasi, nip = :nip
		WHERE no_rawat = :no_rawat
		AND tgl_perawatan = :tgl_perawatan
		AND CAST(jam_rawat AS time(0)) = CAST(:jam_rawat AS time(0))
	`

	log.Printf("[INFO] Attempting update for no_rawat=%s, tgl_perawatan=%s, jam_rawat=%s", p.NoRawat, p.TglPerawatan, p.JamRawat)
	log.Printf("[DEBUG] suhu: %s, tensi: %s, keluhan: %s", p.SuhuTubuh, p.Tensi, p.Keluhan)
	res, err := r.DB.NamedExec(query, p)
	if err != nil {
		log.Printf("[ERROR] Failed to execute update: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	log.Printf("[DEBUG] Rows affected: %d", rowsAffected)
	if err != nil {
		log.Printf("[ERROR] Could not get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("[WARN] No rows updated for no_rawat=%s, tgl_perawatan=%s, jam_rawat=%s", p.NoRawat, p.TglPerawatan, p.JamRawat)
		return fmt.Errorf("no record updated — no match for no_rawat=%s, tgl_perawatan=%s, jam_rawat=%s", p.NoRawat, p.TglPerawatan, p.JamRawat)
	}

	log.Printf("[SUCCESS] Updated pemeriksaan_ranap row for no_rawat=%s", p.NoRawat)
	return nil
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
