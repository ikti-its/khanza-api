package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/dokter/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.Dokter, error)
	FindById(id string) (entity.Dokter, error)
	Insert(data *entity.Dokter) error
	Update(data *entity.Dokter) error
	Delete(id string) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.Dokter, error) {
	query := `SELECT * FROM dokter ORDER BY kode_dokter ASC`
	var result []entity.Dokter
	err := r.DB.Select(&result, query)
	return result, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Dokter, error) {
	query := `SELECT * FROM dokter WHERE kode_dokter = $1`
	var dokter entity.Dokter
	err := r.DB.Get(&dokter, query, id)
	return dokter, err
}

func (r *RepositoryImpl) Insert(data *entity.Dokter) error {
	query := `
	INSERT INTO dokter (
		kode_dokter, nama_dokter, jenis_kelamin, alamat_tinggal,
		no_telp, spesialis, izin_praktik
	) VALUES (
		$1, $2, $3, $4,
		$5, $6, $7
	)`
	_, err := r.DB.Exec(query,
		data.Kode_dokter,
		data.Nama_dokter,
		data.Jenis_kelamin,
		data.Alamat_tinggal,
		data.No_telp,
		data.Spesialis,
		data.Izin_praktik,
	)
	return err
}

func (r *RepositoryImpl) Update(data *entity.Dokter) error {
	query := `
	UPDATE dokter SET 
		nama_dokter = $2, jenis_kelamin = $3, alamat_tinggal = $4,
		no_telp = $5, spesialis = $6, izin_praktik = $7
	WHERE kode_dokter = $1
	`
	_, err := r.DB.Exec(query,
		data.Kode_dokter,
		data.Nama_dokter,
		data.Jenis_kelamin,
		data.Alamat_tinggal,
		data.No_telp,
		data.Spesialis,
		data.Izin_praktik,
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `DELETE FROM dokter WHERE kode_dokter = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
