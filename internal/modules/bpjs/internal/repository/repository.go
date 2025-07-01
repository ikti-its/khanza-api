package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/bpjs/internal/entity"
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
		SELECT * FROM bpjs ORDER BY no_bpjs DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM bpjs WHERE no_bpjs = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO bpjs (
			no_bpjs, nama_program, penyelenggara, tarif, batas_bawah, batas_atas
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_bpjs,    
		entity.Nama_program,   
		entity.Penyelenggara,   
		entity.Tarif,       
		entity.Batas_bawah,  
		entity.Batas_atas, 
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE bpjs SET 
			nama_program = $2, penyelenggara = $3, tarif = $4, batas_bawah = $5, batas_atas = $6
		WHERE no_bpjs = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_bpjs,    
		entity.Nama_program,   
		entity.Penyelenggara,   
		entity.Tarif,       
		entity.Batas_bawah,  
		entity.Batas_atas, 
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM bpjs WHERE no_bpjs = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
