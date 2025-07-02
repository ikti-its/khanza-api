package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/thr/internal/entity"
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
		SELECT * FROM thr ORDER BY no_thr DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM thr WHERE no_thr = $1`

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
		INSERT INTO thr (
			no_thr, masa_kerja, pengali_upah
		) VALUES (
			$1, $2, $3
		)
	`
	_, err := r.DB.Exec(query,
		entity.No_thr,    
		entity.Masa_kerja,   
		entity.Pengali_upah,   
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
		UPDATE thr SET 
			masa_kerja = $2, pengali_upah = $3
		WHERE no_thr = $1
	`
	_, err := r.DB.Exec(query,
		entity.No_thr,    
		entity.Masa_kerja,   
		entity.Pengali_upah,  
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
		DELETE FROM thr WHERE no_thr = $1
	`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
