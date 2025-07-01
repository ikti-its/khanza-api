package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/pph21/internal/entity"
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
		SELECT * FROM pph21 ORDER BY no_pph21 DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM pph21 WHERE no_pph21 = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO pph21 (
			no_pph21, pkp_bawah, pkp_atas, tarif_pajak
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_pph21,  
		entity.Pkp_bawah,
		entity.Pkp_atas,  
		entity.Tarif_pajak,   
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE pph21 SET 
			pkp_bawah = $2, pkp_atas = $3, tarif_pajak = $4
		WHERE no_pph21 = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_pph21,  
		entity.Pkp_bawah,
		entity.Pkp_atas,  
		entity.Tarif_pajak,   
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM pph21 WHERE no_pph21 = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
