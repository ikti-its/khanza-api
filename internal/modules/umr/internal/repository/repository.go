package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/umr/internal/entity"
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
		SELECT * FROM umr ORDER BY no_umr DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM umr WHERE no_umr = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO umr (
			no_umr, provinsi, kotakab, jenis, upah_minimum
		) VALUES (
			$1, $2, $3, $4, $5
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_umr,    
		entity.Provinsi,   
		entity.Kotakab,   
		entity.Jenis,
		entity.Upah_minimum, 
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE umr SET 
			provinsi = $2, kotakab = $3, jenis = $4, upah_minimum = $5
	`
	_, err := r.DB.Exec(query,
		entity.No_umr,    
		entity.Provinsi,   
		entity.Kotakab,   
		entity.Jenis,
		entity.Upah_minimum,
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM umr WHERE no_umr = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
