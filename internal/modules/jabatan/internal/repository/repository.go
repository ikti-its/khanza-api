package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/jabatan/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.Entity, error)
	FindById(id string) (entity.Entity, error)
	Insert(c *fiber.Ctx, entity *entity.Entity) error
	Update(c *fiber.Ctx, entity *entity.Entity) error
	Delete(c *fiber.Ctx, id string) error 
    setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func (r *RepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID, ok := userIDRaw.(string)
	if !ok {
		return fmt.Errorf("invalid user_id type: %T", userIDRaw)
	}
	safeUserID := pq.QuoteLiteral(userID)
	_, err := tx.Exec(fmt.Sprintf(`SET LOCAL my.user_id = %s`, safeUserID))
	return err
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
tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO jabatan_pegawai (
			no_jabatan, jenis_jabatan, nama_jabatan, jenjang, tunjangan
		) VALUES (
			$1, $2, $3, $4, $5
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_jabatan,    
		entity.Jenis_jabatan,   
		entity.Nama_jabatan,   
		entity.Jenjang,       
		entity.Tunjangan,  
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *RepositoryImpl) Update(entity *entity.Entity) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

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
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *RepositoryImpl) Delete(id string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

    query := `
		DELETE FROM jabatan_pegawai WHERE no_jabatan = $1
	`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
