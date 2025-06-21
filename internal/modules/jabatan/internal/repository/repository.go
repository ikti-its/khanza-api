package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/jabatan/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.Entity, error)
	FindById(id string) (entity.Entity, error)
	Insert(entity *entity.Entity) error
	Update(entity *entity.Entity) error
	Delete(id string) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.Entity, error) {
	query := `
		SELECT * FROM jabatan_pegawai ORDER BY no_jabatan DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM jabatan_pegawai WHERE no_jabatan = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO jabatan_pegawai (
			no_jabatan, jenis_jabatan, nama_jabatan, jenjang, tunjangan
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_jabatan,    
		entity.Jenis_jabatan,   
		entity.Nama_jabatan,   
		entity.Jenjang,       
		entity.Tunjangan,  
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE jabatan_pegawai SET 
			jenis_jabatan = $2, nama_jabatan = $3, jenis = $4, tunjangan = $5
		WHERE no_jabatan = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_jabatan,    
		entity.Jenis_jabatan,   
		entity.Nama_jabatan,   
		entity.Jenjang,       
		entity.Tunjangan,    
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM jabatan_pegawai WHERE no_jabatan = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
