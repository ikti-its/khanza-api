package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/upmk/internal/entity"
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
		SELECT * FROM upmk ORDER BY no_upmk DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM upmk WHERE no_upmk = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO upmk (
			no_upmk, masa_kerja, pengali_upah
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_upmk,    
		entity.Masa_kerja,   
		entity.Pengali_upah,   
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE upmk SET 
			masa_kerja = $2, pengali_upah = $3
		WHERE no_upmk = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_upmk,    
		entity.Masa_kerja,   
		entity.Pengali_upah,  
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM upmk WHERE no_upmk = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
