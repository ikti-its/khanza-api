package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/umr/internal/entity"
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
tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

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
		UPDATE umr SET 
			provinsi = $2, kotakab = $3, jenis = $4, upah_minimum = $5
		WHERE no_umr = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_umr,    
		entity.Provinsi,   
		entity.Kotakab,   
		entity.Jenis,
		entity.Upah_minimum,
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
		DELETE FROM umr WHERE no_umr = $1
	`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
