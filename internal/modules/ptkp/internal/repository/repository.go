package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/ptkp/internal/entity"
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
		SELECT * FROM ptkp ORDER BY no_ptkp DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM ptkp WHERE no_ptkp = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO ptkp (
			no_ptkp, kode_ptkp, perkawinan, tanggungan, nilai_ptkp
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_ptkp,    
		entity.Kode_ptkp,   
		entity.Perkawinan,   
		entity.Tanggungan,       
		entity.Nilai_ptkp,  
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE ptkp_pegawai SET 
			kode_ptkp = $2, perkawinan = $3, tanggungan = $4, nilai_ptkp = $5
		WHERE no_ptkp = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_ptkp,    
		entity.Kode_ptkp,   
		entity.Perkawinan,   
		entity.Tanggungan,       
		entity.Nilai_ptkp, 
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM ptkp_pegawai WHERE no_ptkp = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
