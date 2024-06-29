package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type pemesananRepositoryImpl struct {
	DB *sqlx.DB
}

func NewPemesananRepository(db *sqlx.DB) repository.PemesananRepository {
	return &pemesananRepositoryImpl{db}
}

func (r *pemesananRepositoryImpl) Insert(pemesanan *entity.Pemesanan) error {
	query := `
		INSERT INTO pemesanan_barang_medis (id, tanggal_pesan, no_pemesanan, id_pengajuan, id_supplier, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pemesanan, updater)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.DB.Exec(query, pemesanan.Id, pemesanan.Tanggal, pemesanan.Nomor, pemesanan.IdPengajuan, pemesanan.Supplier, pemesanan.IdPegawai, pemesanan.DiskonPersen, pemesanan.DiskonJumlah, pemesanan.PajakPersen, pemesanan.PajakJumlah, pemesanan.Materai, pemesanan.Total, pemesanan.Updater)

	return err
}

func (r *pemesananRepositoryImpl) Find() ([]entity.Pemesanan, error) {
	query := `
		SELECT id, tanggal_pesan, no_pemesanan, id_pengajuan, id_supplier, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pemesanan
		FROM pemesanan_barang_medis
		WHERE deleted_at IS NULL
	`

	var records []entity.Pemesanan
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *pemesananRepositoryImpl) FindPage(page, size int) ([]entity.Pemesanan, int, error) {
	query := `
		SELECT id, tanggal_pesan, no_pemesanan, id_pengajuan, id_supplier, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pemesanan
		FROM pemesanan_barang_medis
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM pemesanan_barang_medis WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Pemesanan
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *pemesananRepositoryImpl) FindById(id uuid.UUID) (entity.Pemesanan, error) {
	query := `
		SELECT id, tanggal_pesan, no_pemesanan, id_pengajuan, id_supplier, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pemesanan
		FROM pemesanan_barang_medis
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Pemesanan
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *pemesananRepositoryImpl) Update(pemesanan *entity.Pemesanan) error {
	query := `
		UPDATE pemesanan_barang_medis
		SET tanggal_pesan = $2, no_pemesanan = $3, id_pengajuan = $4, id_supplier = $5, id_pegawai = $6, diskon_persen = $7, diskon_jumlah = $8, pajak_persen = $9, pajak_jumlah = $10, materai = $11, total_pemesanan = $14, updated_at = $12, updater = $13
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, pemesanan.Id, pemesanan.Tanggal, pemesanan.Nomor, pemesanan.IdPengajuan, pemesanan.Supplier, pemesanan.IdPegawai, pemesanan.DiskonPersen, pemesanan.DiskonJumlah, pemesanan.PajakPersen, pemesanan.PajakJumlah, pemesanan.Materai, time.Now(), pemesanan.Updater, pemesanan.Total)

	return err
}

func (r *pemesananRepositoryImpl) Delete(pemesanan *entity.Pemesanan) error {
	query := `
		UPDATE pemesanan_barang_medis
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, pemesanan.Id, time.Now(), pemesanan.Updater)

	return err
}
