package postgres

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type diagnosaPasienRepositoryImpl struct {
	DB *sqlx.DB
}

func NewDiagnosaPasienRepository(db *sqlx.DB) repository.DiagnosaPasienRepository {
	return &diagnosaPasienRepositoryImpl{DB: db}
}

func (r *diagnosaPasienRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID, ok := userIDRaw.(string)
	if !ok {
		log.Println("⚠️ user_id is not a string")
		return fmt.Errorf("invalid user_id type: expected string, got %T", userIDRaw)
	}

	safeUserID := pq.QuoteLiteral(userID)
	query := fmt.Sprintf(`SET LOCAL my.user_id = %s`, safeUserID)

	_, err := tx.Exec(query)
	if err != nil {
		log.Printf("❌ Failed to SET LOCAL my.user_id = %v: %v\n", userID, err)
	}
	return err
}

func (r *diagnosaPasienRepositoryImpl) Insert(c *fiber.Ctx, data *entity.DiagnosaPasien) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO diagnosa_pasien (
			no_rawat, kd_penyakit, status, prioritas, status_penyakit
		) VALUES (
			:no_rawat, :kd_penyakit, :status, :prioritas, :status_penyakit
		)`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *diagnosaPasienRepositoryImpl) FindAll() ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien ORDER BY no_rawat, prioritas ASC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien WHERE no_rawat = $1 ORDER BY prioritas ASC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) FindByKodePenyakit(kode string) ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien WHERE kd_penyakit = $1`
	err := r.DB.Select(&list, query, kode)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) FindByNoRawatAndStatus(noRawat string, status string) ([]entity.DiagnosaPasien, error) {
	var list []entity.DiagnosaPasien
	query := `SELECT * FROM diagnosa_pasien WHERE no_rawat = $1 AND status = $2 ORDER BY prioritas ASC`
	err := r.DB.Select(&list, query, noRawat, status)
	return list, err
}

func (r *diagnosaPasienRepositoryImpl) Update(c *fiber.Ctx, data *entity.DiagnosaPasien) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		UPDATE diagnosa_pasien SET
			status = :status,
			prioritas = :prioritas,
			status_penyakit = :status_penyakit
		WHERE no_rawat = :no_rawat AND kd_penyakit = :kd_penyakit
	`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *diagnosaPasienRepositoryImpl) Delete(c *fiber.Ctx, noRawat string, kdPenyakit string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `DELETE FROM diagnosa_pasien WHERE no_rawat = $1 AND kd_penyakit = $2`
	_, err = tx.Exec(query, noRawat, kdPenyakit)
	if err != nil {
		return err
	}

	return tx.Commit()
}
