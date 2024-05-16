package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type obatRepositoryImpl struct {
	DB *sqlx.DB
}

func NewObatRepository(db *sqlx.DB) repository.ObatRepository {
	return &obatRepositoryImpl{db}
}

func (r *obatRepositoryImpl) Insert(obat *entity.Obat) error {
	query := `
		INSERT INTO obat (id, id_barang_medis, id_industri_farmasi, kandungan, id_satuan_kecil, isi, kapasitas, id_jenis, id_kategori, id_golongan, kadaluwarsa, updater)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.DB.Exec(query, obat.Id, obat.IdMedis, obat.Industri, obat.Kandungan, obat.Satuan, obat.Isi, obat.Kapasitas, obat.Jenis, obat.Kategori, obat.Golongan, obat.Kadaluwarsa, obat.Updater)

	return err
}

func (r *obatRepositoryImpl) Find() ([]entity.Obat, error) {
	query := `
		SELECT id, id_barang_medis, id_industri_farmasi, kandungan, id_satuan_kecil, isi, kapasitas, id_jenis, id_kategori, id_golongan, kadaluwarsa
		FROM obat
		WHERE deleted_at IS NULL
	`

	var records []entity.Obat
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *obatRepositoryImpl) FindPage(page, size int) ([]entity.Obat, int, error) {
	query := `
		SELECT id, id_barang_medis, id_industri_farmasi, kandungan, id_satuan_kecil, isi, kapasitas, id_jenis, id_kategori, id_golongan, kadaluwarsa
		FROM obat
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM obat WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Obat
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *obatRepositoryImpl) FindById(id uuid.UUID) (entity.Obat, error) {
	query := `
		SELECT id, id_barang_medis, id_industri_farmasi, kandungan, id_satuan_kecil, isi, kapasitas, id_jenis, id_kategori, id_golongan, kadaluwarsa
		FROM obat
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Obat
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *obatRepositoryImpl) FindByIdMedis(id uuid.UUID) (entity.Obat, error) {
	query := `
		SELECT id, id_barang_medis, id_industri_farmasi, kandungan, id_satuan_kecil, isi, kapasitas, id_jenis, id_kategori, id_golongan, kadaluwarsa
		FROM obat
		WHERE id_barang_medis = $1 AND deleted_at IS NULL
	`

	var record entity.Obat
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *obatRepositoryImpl) Update(obat *entity.Obat) error {
	query := `
		UPDATE obat 
		SET id_barang_medis = $2, id_industri_farmasi = $3, kandungan = $4, id_satuan_kecil = $5, isi = $6, kapasitas = $7, id_jenis = $8, id_kategori = $9, id_golongan = $10, kadaluwarsa = $11, updated_at = $12, updater = $13
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, obat.Id, obat.IdMedis, obat.Industri, obat.Kandungan, obat.Satuan, obat.Isi, obat.Kapasitas, obat.Jenis, obat.Kategori, obat.Golongan, obat.Kadaluwarsa, time.Now(), obat.Updater)

	return err
}

func (r *obatRepositoryImpl) Delete(obat *entity.Obat) error {
	query := `
		UPDATE obat 
		SET deleted_at = $1, updater = $2
		WHERE id = $3
	`

	_, err := r.DB.Exec(query, time.Now(), obat.Updater, obat.Id)

	return err
}
