package postgres

import (
	"fmt"

	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/registrasi/internal/repository"
	"github.com/jmoiron/sqlx"
)

type RegistrasiRepository interface {
	Insert(registrasi *entity.Registrasi) error
	Find() ([]entity.Registrasi, error)
	FindAll() ([]entity.Registrasi, error)
	FindByNomorReg(nomorReg string) (entity.Registrasi, error)
	FindByNomorRM(nomorReg string) (entity.Registrasi, error)
	FindByTanggal(nomorReg string) (entity.Registrasi, error)
	Update(registrasi *entity.Registrasi) error
	Delete(nomorReg string) error
	UpdateStatusKamar(nomorReg string, status string) error
	AssignKamar(nomorReg string, kamarID string) error

	CheckDokterExists(kodeDokter string) (bool, error)
}

type registrasiRepositoryImpl struct {
	DB *sqlx.DB
}

func (r *registrasiRepositoryImpl) UpdateStatusKamar(nomorReg, status string) error {
	query := `UPDATE registrasi SET status_kamar = $1 WHERE nomor_reg = $2`
	_, err := r.DB.Exec(query, status, nomorReg)
	return err
}

func (r *registrasiRepositoryImpl) CheckDokterExists(kodeDokter string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM dokter WHERE kode_dokter = $1)`

	err := r.DB.QueryRow(query, kodeDokter).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewRegistrasiRepository(db *sqlx.DB) repository.RegistrasiRepository {
	return &registrasiRepositoryImpl{DB: db}
}

func (r *registrasiRepositoryImpl) Insert(registrasi *entity.Registrasi) error {
	query := `
		INSERT INTO registrasi (
			nomor_reg, nomor_rawat, tanggal, jam, kode_dokter, nama_dokter, nomor_rm,
			nama_pasien, jenis_kelamin, umur, poliklinik, jenis_bayar, penanggung_jawab,
			alamat_pj, hubungan_pj, biaya_registrasi, status_registrasi, no_telepon,
			status_rawat, status_poli, status_bayar, status_kamar
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18,
			$19, $20, $21, $22
		)
	`
	_, err := r.DB.Exec(query,
		registrasi.NomorReg, registrasi.NomorRawat, registrasi.Tanggal, registrasi.Jam,
		registrasi.KodeDokter, registrasi.NamaDokter, registrasi.NomorRM, registrasi.Nama,
		registrasi.JenisKelamin, registrasi.Umur, registrasi.Poliklinik, registrasi.JenisBayar,
		registrasi.PenanggungJawab, registrasi.Alamat, registrasi.HubunganPJ, registrasi.BiayaRegistrasi,
		registrasi.StatusRegistrasi, registrasi.NoTelepon, registrasi.StatusRawat,
		registrasi.StatusPoli, registrasi.StatusBayar, registrasi.StatusKamar, // ✅ Added
	)
	return err
}

func (r *registrasiRepositoryImpl) Find() ([]entity.Registrasi, error) {
	query := `
		SELECT * FROM registrasi ORDER BY tanggal DESC
	`
	var records []entity.Registrasi
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *registrasiRepositoryImpl) FindAll() ([]entity.Registrasi, error) {
	query := `
		SELECT * FROM registrasi ORDER BY tanggal DESC
	`
	var records []entity.Registrasi
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *registrasiRepositoryImpl) FindByNomorReg(nomorReg string) (entity.Registrasi, error) {
	query := `
		SELECT * FROM registrasi WHERE nomor_reg = $1
	`
	var record entity.Registrasi
	err := r.DB.Get(&record, query, nomorReg)
	return record, err
}

func (r *registrasiRepositoryImpl) FindByNomorRM(nomorReg string) (entity.Registrasi, error) {
	query := `
		SELECT * FROM registrasi WHERE nomor_rm = $1
	`
	var record entity.Registrasi
	err := r.DB.Get(&record, query, nomorReg)
	return record, err
}

func (r *registrasiRepositoryImpl) FindByTanggal(nomorReg string) (entity.Registrasi, error) {
	query := `
		SELECT * FROM registrasi WHERE nomor_rm = $1
	`
	var record entity.Registrasi
	err := r.DB.Get(&record, query, nomorReg)
	return record, err
}

func (r *registrasiRepositoryImpl) Update(registrasi *entity.Registrasi) error {
	query := `
		UPDATE registrasi SET 
			nomor_rawat = $2, tanggal = $3, jam = $4, kode_dokter = $5, nama_dokter = $6,
			nomor_rm = $7, nama_pasien = $8, jenis_kelamin = $9, umur = $10, poliklinik = $11,
			jenis_bayar = $12, penanggung_jawab = $13, alamat_pj = $14, hubungan_pj = $15,
			biaya_registrasi = $16, status_registrasi = $17, no_telepon = $18,
			status_rawat = $19, status_poli = $20, status_bayar = $21, status_kamar = $22
		WHERE nomor_reg = $1
	`
	_, err := r.DB.Exec(query,
		registrasi.NomorReg, registrasi.NomorRawat, registrasi.Tanggal, registrasi.Jam,
		registrasi.KodeDokter, registrasi.NamaDokter, registrasi.NomorRM, registrasi.Nama,
		registrasi.JenisKelamin, registrasi.Umur, registrasi.Poliklinik, registrasi.JenisBayar,
		registrasi.PenanggungJawab, registrasi.Alamat, registrasi.HubunganPJ, registrasi.BiayaRegistrasi,
		registrasi.StatusRegistrasi, registrasi.NoTelepon, registrasi.StatusRawat,
		registrasi.StatusPoli, registrasi.StatusBayar, registrasi.StatusKamar, // ✅ Added
	)
	return err
}

func (r *registrasiRepositoryImpl) Delete(nomorReg string) error {
	query := `
		DELETE FROM registrasi WHERE nomor_reg = $1
	`
	_, err := r.DB.Exec(query, nomorReg)
	return err
}

func (r *registrasiRepositoryImpl) FindPendingRoomRequests() ([]entity.Registrasi, error) {
	query := `
		SELECT nomor_reg, nama_pasien, nomor_rm, status_kamar
		FROM registrasi
		WHERE status_kamar = 'menunggu'
		`
	var results []entity.Registrasi

	fmt.Println("✅ Running pending-room query...")

	err := r.DB.Select(&results, query)
	fmt.Printf("🔍 Query returned %d rows\n", len(results))
	if err != nil {
		fmt.Println("❌ DB error:", err)
		for _, r := range results {
			fmt.Printf("🔹 %s | %s | %s | %s\n", r.NomorReg, r.Nama, r.NomorRM, r.StatusKamar)
		}
	}
	return results, err
}

func (r *registrasiRepositoryImpl) AssignKamar(nomorReg string, nomorBed string) error {
	// 1. Update registrasi
	updateQuery := `UPDATE registrasi SET nomor_bed = $1, status_kamar = 'diterima' WHERE nomor_reg = $2`
	_, err := r.DB.Exec(updateQuery, nomorBed, nomorReg)
	if err != nil {
		return err
	}

	// 2. Insert into rawatinap
	insertQuery := `
	INSERT INTO rawat_inap (
		nomor_rawat, nomor_rm, nama_pasien, alamat_pasien, penanggung_jawab, hubungan_pj,
		jenis_bayar, kamar, tarif_kamar, tanggal_masuk, jam_masuk, dokter_pj, diagnosa_awal, status_pulang
	)
	SELECT 
		r.nomor_rawat, r.nomor_rm, r.nama_pasien, r.alamat_pj, r.penanggung_jawab, r.hubungan_pj,
		r.jenis_bayar, k.nomor_bed, k.tarif_kamar, CURRENT_DATE, CURRENT_TIME,
		r.nama_dokter, '', 'Belum'
	FROM registrasi r
	JOIN kamar k ON k.nomor_bed = $1
	WHERE r.nomor_reg = $2
	`
	_, err = r.DB.Exec(insertQuery, nomorBed, nomorReg)
	return err
}
