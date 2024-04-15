package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/akun/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/akun/internal/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

type alamatRepositoryImpl struct {
	DB *sqlx.DB
}

func NewAlamatRepository(db *sqlx.DB) repository.AlamatRepository {
	return &alamatRepositoryImpl{db}
}

func (r *alamatRepositoryImpl) Insert(alamat *entity.Alamat) error {
	query := `
		INSERT INTO alamat (id_akun, alamat, alamat_lat, alamat_lon, kota, kode_pos) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.DB.Exec(query, alamat.IdAkun, alamat.Alamat, alamat.AlamatLat, alamat.AlamatLon, alamat.Kota, alamat.KodePos)

	return err
}

func (r *alamatRepositoryImpl) FindById(id uuid.UUID) (entity.Alamat, error) {
	query := `
		SELECT id_akun, alamat, alamat_lat, alamat_lon, kota, kode_pos 
		FROM alamat 
		WHERE id_akun = $1 AND deleted_at IS NULL
	`

	var record entity.Alamat
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *alamatRepositoryImpl) Update(alamat *entity.Alamat) error {
	query := `
		UPDATE alamat 
		SET id_akun = $1, alamat = $2, alamat_lat = $3, alamat_lon = $4, kota = $5, kode_pos = $6, updated_at = $7, updater = $8 
		WHERE id_akun = $9 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, alamat.IdAkun, alamat.Alamat, alamat.AlamatLat, alamat.AlamatLon, alamat.Kota, alamat.KodePos, time.Now(), alamat.Updater, alamat.IdAkun)

	return err
}

func (r *alamatRepositoryImpl) Delete(alamat *entity.Alamat) error {
	query := `
		UPDATE alamat 
		SET deleted_at = $1, updater = $2
		WHERE id_akun = $3
	`

	_, err := r.DB.Exec(query, time.Now(), alamat.Updater, alamat.IdAkun)

	return err
}
