package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type alkesRepositoryImpl struct {
	DB *sqlx.DB
}

func NewAlkesRepository(db *sqlx.DB) repository.AlkesRepository {
	return &alkesRepositoryImpl{db}
}

func (r *alkesRepositoryImpl) Insert(alkes *entity.Alkes) error {
	query := `
		INSERT INTO alat_kesehatan (id, id_barang_medis, merek, updater)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.DB.Exec(query, alkes.Id, alkes.IdMedis, alkes.Merek, alkes.Updater)

	return err
}

func (r *alkesRepositoryImpl) Find() ([]entity.Alkes, error) {
	query := `
		SELECT id, id_barang_medis, merek
		FROM alat_kesehatan
		WHERE deleted_at IS NULL
	`

	var records []entity.Alkes
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *alkesRepositoryImpl) FindPage(page, size int) ([]entity.Alkes, int, error) {
	query := `
		SELECT id, id_barang_medis, merek
		FROM alat_kesehatan
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM alat_kesehatan WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Alkes
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *alkesRepositoryImpl) FindById(id uuid.UUID) (entity.Alkes, error) {
	query := `
		SELECT id, id_barang_medis, merek
		FROM alat_kesehatan
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Alkes
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *alkesRepositoryImpl) FindByIdMedis(id uuid.UUID) (entity.Alkes, error) {
	query := `
		SELECT id, id_barang_medis, merek
		FROM alat_kesehatan
		WHERE id_barang_medis = $1 AND deleted_at IS NULL
	`

	var record entity.Alkes
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *alkesRepositoryImpl) Update(alkes *entity.Alkes) error {
	query := `
		UPDATE alat_kesehatan
		SET id_barang_medis = $2, merek = $3, updated_at = $4, updater = $5
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, alkes.Id, alkes.IdMedis, alkes.Merek, time.Now(), alkes.Updater)

	return err
}

func (r *alkesRepositoryImpl) Delete(alkes *entity.Alkes) error {
	query := `
		UPDATE barang_medis
		SET deleted_at = $1, updater = $2
		WHERE id = $3
	`

	_, err := r.DB.Exec(query, time.Now(), alkes.Updater, alkes.Id)

	return err
}
