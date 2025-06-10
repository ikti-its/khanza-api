package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/rawatinap/internal/entity"
	"github.com/jmoiron/sqlx"
)

type RawatInapRepository interface {
	Insert(rawatInap *entity.RawatInap) error
	FindAll() ([]entity.RawatInap, error)
	FindByNomorRawat(nomorRawat string) (entity.RawatInap, error)
	FindByNomorRM(nomorRM string) ([]entity.RawatInap, error)
	FindByTanggalMasuk(tanggal string) ([]entity.RawatInap, error)
	Update(rawatInap *entity.RawatInap) error
	Delete(nomorRawat string) error
}

type rawatInapRepositoryImpl struct {
	DB *sqlx.DB
}

func NewRawatInapRepository(db *sqlx.DB) RawatInapRepository {
	return &rawatInapRepositoryImpl{DB: db}
}

func (r *rawatInapRepositoryImpl) Insert(ri *entity.RawatInap) error {
	query := `
		INSERT INTO rawat_inap (
			nomor_rawat, nomor_rm, nama_pasien, alamat_pasien, penanggung_jawab,
			hubungan_pj, jenis_bayar, kamar, tarif_kamar, diagnosa_awal, diagnosa_akhir,
			tanggal_masuk, jam_masuk, tanggal_keluar, jam_keluar, total_biaya,
			status_pulang, lama_ranap, dokter_pj, status_bayar
		) VALUES (
			:nomor_rawat, :nomor_rm, :nama_pasien, :alamat_pasien, :penanggung_jawab,
			:hubungan_pj, :jenis_bayar, :kamar, :tarif_kamar, :diagnosa_awal, :diagnosa_akhir,
			:tanggal_masuk, :jam_masuk, :tanggal_keluar, :jam_keluar, :total_biaya,
			:status_pulang, :lama_ranap, :dokter_pj, :status_bayar
		)
	`
	_, err := r.DB.NamedExec(query, ri)
	return err
}

func (r *rawatInapRepositoryImpl) FindAll() ([]entity.RawatInap, error) {
	query := `SELECT * FROM rawat_inap ORDER BY tanggal_masuk DESC`
	var result []entity.RawatInap
	err := r.DB.Select(&result, query)
	return result, err
}

func (r *rawatInapRepositoryImpl) FindByNomorRawat(nomorRawat string) (entity.RawatInap, error) {
	query := `SELECT * FROM rawat_inap WHERE nomor_rawat = $1`
	var result entity.RawatInap
	err := r.DB.Get(&result, query, nomorRawat)
	return result, err
}

func (r *rawatInapRepositoryImpl) FindByNomorRM(nomorRM string) ([]entity.RawatInap, error) {
	query := `SELECT * FROM rawat_inap WHERE nomor_rm = $1 ORDER BY tanggal_masuk DESC`
	var result []entity.RawatInap
	err := r.DB.Select(&result, query, nomorRM)
	return result, err
}

func (r *rawatInapRepositoryImpl) FindByTanggalMasuk(tanggal string) ([]entity.RawatInap, error) {
	query := `SELECT * FROM rawat_inap WHERE tanggal_masuk = $1 ORDER BY jam_masuk`
	var result []entity.RawatInap
	err := r.DB.Select(&result, query, tanggal)
	return result, err
}

func (r *rawatInapRepositoryImpl) Update(ri *entity.RawatInap) error {
	query := `
		UPDATE rawat_inap SET
			nomor_rm = :nomor_rm,
			nama_pasien = :nama_pasien,
			alamat_pasien = :alamat_pasien,
			penanggung_jawab = :penanggung_jawab,
			hubungan_pj = :hubungan_pj,
			jenis_bayar = :jenis_bayar,
			kamar = :kamar,
			tarif_kamar = :tarif_kamar,
			diagnosa_awal = :diagnosa_awal,
			diagnosa_akhir = :diagnosa_akhir,
			tanggal_masuk = :tanggal_masuk,
			jam_masuk = :jam_masuk,
			tanggal_keluar = :tanggal_keluar,
			jam_keluar = :jam_keluar,
			total_biaya = :total_biaya,
			status_pulang = :status_pulang,
			lama_ranap = :lama_ranap,
			dokter_pj = :dokter_pj,
			status_bayar = :status_bayar
		WHERE nomor_rawat = :nomor_rawat
	`
	_, err := r.DB.NamedExec(query, ri)
	return err
}

func (r *rawatInapRepositoryImpl) Delete(nomorRawat string) error {
	query := `DELETE FROM rawat_inap WHERE nomor_rawat = $1`
	_, err := r.DB.Exec(query, nomorRawat)
	return err
}
