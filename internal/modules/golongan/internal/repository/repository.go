package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/golongan/internal/entity"
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
		SELECT * FROM golongan ORDER BY no_golongan DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM golongan WHERE no_golongan = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(entity *entity.Entity) error {
	query := `
		INSERT INTO golongan (
			no_golongan, kode_golongan, nama_golongan, pendidikan, gaji_pokok
		) VALUES (
			$1, $2, $3, $4, $5
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_golongan,    
		entity.Kode_golongan,   
		entity.Nama_golongan,   
		entity.Pendidikan,       
		entity.Gaji_pokok,  
	)
	return err
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	query := `
		UPDATE golongan SET 
			kode_golongan = $2, nama_golongan = $3, pendidikan = $4, gaji_pokok = $5
		WHERE no_golongan = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_golongan,    
		entity.Kode_golongan,   
		entity.Nama_golongan,   
		entity.Pendidikan,       
		entity.Gaji_pokok,  
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `
		DELETE FROM golongan WHERE no_golongan = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
