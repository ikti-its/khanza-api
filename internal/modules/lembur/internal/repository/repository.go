package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/lembur/internal/entity"
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
		SELECT * FROM lembur ORDER BY no_lembur DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM lembur WHERE no_lembur = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO lembur (
			no_lembur, jenis_lembur, jam_lembur, pengali_upah
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_lembur,    
		entity.Jenis_lembur,   
		entity.Jam_lembur,   
		entity.Pengali_upah,       
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE lembur SET 
			jenis_lembur = $2, jam_lembur = $3, pengali_upah = $4
		WHERE no_lembur = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_lembur,    
		entity.Jenis_lembur,   
		entity.Jam_lembur,   
		entity.Pengali_upah,  
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM lembur WHERE no_lembur = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
