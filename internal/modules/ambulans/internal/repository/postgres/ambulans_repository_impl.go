package postgres

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AmbulansRepository interface {
	Insert(ambulans *entity.Ambulans) error
	Find() ([]entity.Ambulans, error)
	FindAll() ([]entity.Ambulans, error)
	FindByNoAmbulans(noAmbulans string) (entity.Ambulans, error)
	Update(ambulans *entity.Ambulans) error
	Delete(noAmbulans string) error
	InsertAmbulansRequest(noAmbulans string) error
	FindPendingRequests() ([]entity.Ambulans, error)
	UpdateAmbulansStatus(noAmbulans string, newStatus string) error
	SetPending(noAmbulans string) error
	InsertWithContext(c *fiber.Ctx, ambulans *entity.Ambulans) error
}

type ambulansRepositoryImpl struct {
	DB *sqlx.DB
}

func NewAmbulansRepository(db *sqlx.DB) AmbulansRepository {
	return &ambulansRepositoryImpl{DB: db}
}

func (r *ambulansRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID, ok := userIDRaw.(string)
	if !ok {
		log.Println("⚠️ user_id is not a string")
		return fmt.Errorf("invalid user_id type: expected string, got %T", userIDRaw)
	}

	// ✅ Escape userID safely for SQL string context
	safeUserID := pq.QuoteLiteral(userID) // e.g., turns abc'def -> 'abc''def'

	query := fmt.Sprintf(`SET LOCAL my.user_id = %s`, safeUserID)

	_, err := tx.Exec(query)
	if err != nil {
		log.Printf("❌ Failed to SET LOCAL my.user_id = %v: %v\n", userID, err)
	}
	return err
}

func (r *ambulansRepositoryImpl) InsertWithContext(c *fiber.Ctx, ambulans *entity.Ambulans) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO sik.ambulans (
			no_ambulans, status, supir
		) VALUES (
			$1, $2, $3
		)
	`
	_, err = tx.Exec(query, ambulans.NoAmbulans, ambulans.Status, ambulans.Supir)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ambulansRepositoryImpl) Insert(ambulans *entity.Ambulans) error {
	query := `
		INSERT INTO ambulans (
			no_ambulans, status, supir
		) VALUES (
			$1, $2, $3
		)
	`
	_, err := r.DB.Exec(query,
		ambulans.NoAmbulans,
		ambulans.Status,
		ambulans.Supir,
	)
	return err
}

func (r *ambulansRepositoryImpl) Find() ([]entity.Ambulans, error) {
	query := `SELECT * FROM ambulans ORDER BY no_ambulans DESC`
	var records []entity.Ambulans
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *ambulansRepositoryImpl) FindAll() ([]entity.Ambulans, error) {
	query := `SELECT * FROM ambulans ORDER BY no_ambulans DESC`
	var records []entity.Ambulans
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *ambulansRepositoryImpl) FindByNoAmbulans(noAmbulans string) (entity.Ambulans, error) {
	query := `SELECT * FROM ambulans WHERE no_ambulans = $1`
	var record entity.Ambulans
	err := r.DB.Get(&record, query, noAmbulans)
	return record, err
}

func (r *ambulansRepositoryImpl) Update(ambulans *entity.Ambulans) error {
	query := `
		UPDATE ambulans SET 
			status = $2, supir = $3
		WHERE no_ambulans = $1
	`
	_, err := r.DB.Exec(query,
		ambulans.NoAmbulans,
		ambulans.Status,
		ambulans.Supir,
	)
	return err
}

func (r *ambulansRepositoryImpl) Delete(noAmbulans string) error {
	query := `DELETE FROM ambulans WHERE no_ambulans = $1`
	_, err := r.DB.Exec(query, noAmbulans)
	return err
}

func (r *ambulansRepositoryImpl) InsertAmbulansRequest(noAmbulans string) error {
	query := `INSERT INTO ambulans (no_ambulans, status) VALUES ($1, 'available')`
	_, err := r.DB.Exec(query, noAmbulans)
	return err
}

func (r *ambulansRepositoryImpl) FindPendingRequests() ([]entity.Ambulans, error) {
	query := `SELECT * FROM ambulans`
	var records []entity.Ambulans
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *ambulansRepositoryImpl) UpdateAmbulansStatusWithContext(c *fiber.Ctx, noAmbulans string, newStatus string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `UPDATE ambulans SET status = $1 WHERE no_ambulans = $2`
	_, err = tx.Exec(query, newStatus, noAmbulans)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ambulansRepositoryImpl) UpdateAmbulansStatus(noAmbulans string, newStatus string) error {
	query := `UPDATE ambulans SET status = $1 WHERE no_ambulans = $2`
	_, err := r.DB.Exec(query, newStatus, noAmbulans)
	return err
}

func (r *ambulansRepositoryImpl) SetPending(noAmbulans string) error {
	query := `UPDATE ambulans SET status = 'pending' WHERE no_ambulans = $1`
	_, err := r.DB.Exec(query, noAmbulans)
	return err
}

func (r *ambulansRepositoryImpl) UpdateStatus(noAmbulans string, status string) (int64, error) {
	result, err := r.DB.Exec(`UPDATE ambulans SET status = $1 WHERE no_ambulans = $2`, status, noAmbulans)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
