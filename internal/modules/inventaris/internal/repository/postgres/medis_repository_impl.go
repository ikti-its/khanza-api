package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type medisRepositoryImpl struct {
	DB *sqlx.DB
}

func NewMedisRepository(db *sqlx.DB) repository.MedisRepository {
	return &medisRepositoryImpl{db}
}

func (r *medisRepositoryImpl) Insert(medis *entity.Medis) error {
	query := `
		INSERT INTO barang_medis (id, nama, jenis, harga, stok)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.DB.Exec(query, medis.Id, medis.Nama, medis.Jenis, medis.Harga, medis.Stok)

	return err
}

func (r *medisRepositoryImpl) Find() ([]entity.Medis, error) {
	query := `
		SELECT id, nama, jenis, harga, stok
		FROM barang_medis
		WHERE deleted_at IS NULL
	`

	var records []entity.Medis
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *medisRepositoryImpl) FindPage(page, size int) ([]entity.Medis, int, error) {
	query := `
		SELECT id, nama, jenis, harga, stok
		FROM barang_medis
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM barang_medis WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Medis
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *medisRepositoryImpl) FindByJenis(jenis string) ([]entity.Medis, error) {
	query := `
		SELECT id, nama, jenis, harga, stok
		FROM barang_medis
		WHERE jenis = $1 AND deleted_at IS NULL
	`

	var records []entity.Medis
	err := r.DB.Select(&records, query, jenis)

	return records, err
}

func (r *medisRepositoryImpl) FindById(id uuid.UUID) (entity.Medis, error) {
	query := `
		SELECT id, nama, jenis, harga, stok
		FROM barang_medis
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Medis
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *medisRepositoryImpl) Update(medis *entity.Medis) error {
	query := `
		UPDATE barang_medis
		SET nama = $1, jenis = $2, harga = $3, stok = $4, updated_at = $5, updater = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, medis.Nama, medis.Jenis, medis.Harga, medis.Stok, time.Now(), medis.Updater, medis.Id)

	return err
}

func (r *medisRepositoryImpl) Delete(medis *entity.Medis) error {
	query := `
		UPDATE barang_medis
		SET deleted_at = $1, updater = $2
		WHERE id = $3
	`

	_, err := r.DB.Exec(query, time.Now(), medis.Updater, medis.Id)

	return err
}
