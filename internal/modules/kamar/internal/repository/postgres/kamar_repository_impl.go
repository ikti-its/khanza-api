package postgres

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type KamarRepository interface {
	Find() ([]entity.Kamar, error)
	FindAll() ([]entity.Kamar, error)
	FindByNomorBed(nomorReg string) (entity.Kamar, error)
	FindByKodeKamar(nomorReg string) (entity.Kamar, error)
	Insert(c *fiber.Ctx, kamar *entity.Kamar) error
	Update(c *fiber.Ctx, kamar *entity.Kamar) error
	Delete(c *fiber.Ctx, nomorReg string) error
	GetAvailableRooms() ([]entity.Kamar, error)
	UpdateStatusKamar(nomorBed string, status string) error
	GetDistinctKelas() ([]string, error)
}

type kamarRepositoryImpl struct {
	DB *sqlx.DB
}

func NewKamarRepository(db *sqlx.DB) KamarRepository {
	return &kamarRepositoryImpl{DB: db}
}

func (r *kamarRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID, ok := userIDRaw.(string)
	if !ok {
		return fmt.Errorf("invalid user_id type: %T", userIDRaw)
	}
	safeUserID := pq.QuoteLiteral(userID)
	_, err := tx.Exec(fmt.Sprintf(`SET LOCAL my.user_id = %s`, safeUserID))

	ip_address_Raw := c.Locals("ip_address")
	ip_address, ok2 := ip_address_Raw.(string)
	if !ok2 {
		return fmt.Errorf("invalid ip_address type: %T", ip_address_Raw)
	}
	safe_ip_address := pq.QuoteLiteral(ip_address)
	_, err = tx.Exec(fmt.Sprintf(`SET LOCAL my.ip_address = %s`, safe_ip_address))

	return err
}

func (r *kamarRepositoryImpl) Insert(c *fiber.Ctx, kamar *entity.Kamar) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO kamar (
			nomor_bed, kode_kamar, nama_kamar, kelas, tarif_kamar, status_kamar
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err = tx.Exec(query,
		kamar.NomorBed,
		kamar.KodeKamar,
		kamar.NamaKamar,
		kamar.Kelas,
		kamar.TarifKamar,
		kamar.StatusKamar,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *kamarRepositoryImpl) Find() ([]entity.Kamar, error) {
	query := `
		SELECT * FROM kamar ORDER BY nomor_bed DESC
	`
	var records []entity.Kamar
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *kamarRepositoryImpl) FindAll() ([]entity.Kamar, error) {
	query := `
		SELECT * FROM kamar ORDER BY nomor_bed DESC
	`
	var records []entity.Kamar
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *kamarRepositoryImpl) FindByNomorBed(nomorBed string) (entity.Kamar, error) {
	query := `SELECT * FROM kamar WHERE nomor_bed = $1`

	var record entity.Kamar
	err := r.DB.Get(&record, query, nomorBed)
	return record, err
}

func (r *kamarRepositoryImpl) FindByKodeKamar(nomorBed string) (entity.Kamar, error) {
	query := `
		SELECT * FROM kamar WHERE kode_kamar = $1
	`
	var record entity.Kamar
	err := r.DB.Get(&record, query, nomorBed)
	return record, err
}

func (r *kamarRepositoryImpl) Update(c *fiber.Ctx, kamar *entity.Kamar) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		UPDATE kamar SET 
			kode_kamar = $2, nama_kamar = $3, kelas = $4, tarif_kamar = $5, status_kamar = $6
		WHERE nomor_bed = $1
	`
	_, err = tx.Exec(query,
		kamar.NomorBed,
		kamar.KodeKamar,
		kamar.NamaKamar,
		kamar.Kelas,
		kamar.TarifKamar,
		kamar.StatusKamar,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *kamarRepositoryImpl) Delete(c *fiber.Ctx, nomorReg string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `DELETE FROM kamar WHERE nomor_bed = $1`
	_, err = tx.Exec(query, nomorReg)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *kamarRepositoryImpl) GetAvailableRooms() ([]entity.Kamar, error) {
	query := `SELECT * FROM kamar WHERE status_kamar = 'available'`
	var results []entity.Kamar
	err := r.DB.Select(&results, query)
	return results, err
}

func (r *kamarRepositoryImpl) UpdateStatusKamar(nomorBed, status string) error {
	query := `UPDATE kamar SET status_kamar = $1 WHERE nomor_bed = $2`
	_, err := r.DB.Exec(query, status, nomorBed)
	return err
}

func (r *kamarRepositoryImpl) GetDistinctKelas() ([]string, error) {
	var kelasList []string
	query := "SELECT DISTINCT kelas FROM kamar"

	err := r.DB.Select(&kelasList, query)
	if err != nil {
		return nil, err
	}

	return kelasList, nil
}
