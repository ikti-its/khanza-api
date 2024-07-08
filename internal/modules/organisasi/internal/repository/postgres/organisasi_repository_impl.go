package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type organisasiRepositoryImpl struct {
	DB *sqlx.DB
}

func NewOrganisasiRepository(db *sqlx.DB) repository.OrganisasiRepository {
	return &organisasiRepositoryImpl{db}
}

func (r *organisasiRepositoryImpl) Insert(organisasi *entity.Organisasi) error {
	query := `
		INSERT INTO organisasi (id, nama, alamat, latitude, longitude, radius, updater) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.DB.Exec(query, organisasi.Id, organisasi.Nama, organisasi.Alamat, organisasi.Latitude, organisasi.Longitude, organisasi.Radius, organisasi.Updater)

	return err
}

func (r *organisasiRepositoryImpl) Find() ([]entity.Organisasi, error) {
	query := `
		SELECT id, nama, alamat, latitude, longitude, radius
		FROM organisasi
		WHERE deleted_at IS NULL
	`

	var records []entity.Organisasi
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *organisasiRepositoryImpl) FindPage(page, size int) ([]entity.Organisasi, int, error) {
	query := `
		SELECT id, nama, alamat, latitude, longitude, radius
		FROM organisasi
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM organisasi WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Organisasi
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *organisasiRepositoryImpl) FindCurrent() (entity.Organisasi, error) {
	query := `
		SELECT id, nama, alamat, latitude, longitude, radius
		FROM organisasi
		WHERE deleted_at IS NULL
		LIMIT 1
	`

	var record entity.Organisasi
	err := r.DB.Get(&record, query)

	return record, err
}

func (r *organisasiRepositoryImpl) FindById(id uuid.UUID) (entity.Organisasi, error) {
	query := `
		SELECT id, nama, alamat, latitude, longitude, radius
		FROM organisasi
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Organisasi
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *organisasiRepositoryImpl) Update(organisasi *entity.Organisasi) error {
	query := `
		UPDATE organisasi
		SET nama = $2, alamat = $3, latitude = $4, longitude = $5, radius = $6, updated_at = $7, updater = $8
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, organisasi.Id, organisasi.Nama, organisasi.Alamat, organisasi.Latitude, organisasi.Radius, organisasi.Longitude, time.Now(), organisasi.Updater)

	return err
}

func (r *organisasiRepositoryImpl) Delete(organisasi *entity.Organisasi) error {
	query := `
		UPDATE organisasi 
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, organisasi.Id, time.Now(), organisasi.Updater)

	return err
}
