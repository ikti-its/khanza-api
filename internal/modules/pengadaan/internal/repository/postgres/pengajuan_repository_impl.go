package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type pengajuanRepositoryImpl struct {
	DB *sqlx.DB
}

func NewPengajuanRepository(db *sqlx.DB) repository.PengajuanRepository {
	return &pengajuanRepositoryImpl{db}
}

func (r *pengajuanRepositoryImpl) Insert(pengajuan *entity.Pengajuan) error {
	query := `
		INSERT INTO pengajuan_barang_medis (id, tanggal_pengajuan, nomor_pengajuan, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pengajuan, catatan, status_pesanan, updater)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.DB.Exec(query, pengajuan.Id, pengajuan.Tanggal, pengajuan.Nomor, pengajuan.Pegawai, pengajuan.DiskonPersen, pengajuan.DiskonJumlah, pengajuan.PajakPersen, pengajuan.PajakJumlah, pengajuan.Materai, pengajuan.Total, pengajuan.Catatan, pengajuan.Status, pengajuan.Updater)

	return err
}

func (r *pengajuanRepositoryImpl) Find() ([]entity.Pengajuan, error) {
	query := `
		SELECT id, tanggal_pengajuan, nomor_pengajuan, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pengajuan, catatan, status_pesanan
		FROM pengajuan_barang_medis
		WHERE deleted_at IS NULL
	`

	var records []entity.Pengajuan
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *pengajuanRepositoryImpl) FindPage(page, size int) ([]entity.Pengajuan, int, error) {
	query := `
		SELECT id, tanggal_pengajuan, nomor_pengajuan, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pengajuan, catatan, status_pesanan
		FROM pengajuan_barang_medis
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM pengajuan_barang_medis WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Pengajuan
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *pengajuanRepositoryImpl) FindById(id uuid.UUID) (entity.Pengajuan, error) {
	query := `
		SELECT id, tanggal_pengajuan, nomor_pengajuan, id_pegawai, diskon_persen, diskon_jumlah, pajak_persen, pajak_jumlah, materai, total_pengajuan, catatan, status_pesanan
		FROM pengajuan_barang_medis
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Pengajuan
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *pengajuanRepositoryImpl) Update(pengajuan *entity.Pengajuan) error {
	query := `
		UPDATE pengajuan_barang_medis
		SET tanggal_pengajuan = $2, nomor_pengajuan = $3, id_pegawai = $4, diskon_persen = $5, diskon_jumlah = $6, pajak_persen = $7, pajak_jumlah = $8, materai = $9, total_pengajuan = $14, catatan = $10, status_pesanan = $11, updated_at = $12, updater = $13
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, pengajuan.Id, pengajuan.Tanggal, pengajuan.Nomor, pengajuan.Pegawai, pengajuan.DiskonPersen, pengajuan.DiskonJumlah, pengajuan.PajakPersen, pengajuan.PajakJumlah, pengajuan.Materai, pengajuan.Catatan, pengajuan.Status, time.Now(), pengajuan.Updater, pengajuan.Total)

	return err
}

func (r *pengajuanRepositoryImpl) Delete(pengajuan *entity.Pengajuan) error {
	query := `
		UPDATE pengajuan_barang_medis
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, pengajuan.Id, time.Now(), pengajuan.Updater)

	return err
}
