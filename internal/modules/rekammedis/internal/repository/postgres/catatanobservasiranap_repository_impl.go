package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type catatanObservasiRanapRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCatatanObservasiRanapRepository(db *sqlx.DB) repository.CatatanObservasiRanapRepository {
	return &catatanObservasiRanapRepositoryImpl{DB: db}
}

func (r *catatanObservasiRanapRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID, ok := userIDRaw.(string)
	if !ok {
		log.Println("⚠️ user_id is not a string")
		return fmt.Errorf("invalid user_id type: expected string, got %T", userIDRaw)
	}

	safeUserID := pq.QuoteLiteral(userID) // escapes single quotes safely
	query := fmt.Sprintf(`SET LOCAL my.user_id = %s`, safeUserID)

	if _, err := tx.Exec(query); err != nil {
		log.Printf("❌ Failed to SET LOCAL my.user_id = %v: %v\n", userID, err)
		return err
	}
	return nil
}

func (r *catatanObservasiRanapRepositoryImpl) FindAll() ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap WHERE no_rawat = $1 ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap WHERE tgl_perawatan = $1 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, tanggal)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanap, error) {
	var list []entity.CatatanObservasiRanap
	query := `SELECT * FROM catatan_observasi_ranap WHERE no_rawat = $1 AND tgl_perawatan = $2 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat, tanggal)
	return list, err
}

func (r *catatanObservasiRanapRepositoryImpl) Insert(c *fiber.Ctx, data *entity.CatatanObservasiRanap) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO catatan_observasi_ranap (
			no_rawat, tgl_perawatan, jam_rawat, gcs, td, hr, rr, suhu, spo2, nip
		) VALUES (
			:no_rawat, :tgl_perawatan, :jam_rawat, :gcs, :td, :hr, :rr, :suhu, :spo2, :nip
		)`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *catatanObservasiRanapRepositoryImpl) Update(c *fiber.Ctx, data *entity.CatatanObservasiRanap) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		UPDATE catatan_observasi_ranap SET
			gcs = :gcs, td = :td, hr = :hr, rr = :rr, suhu = :suhu, spo2 = :spo2, nip = :nip
		WHERE no_rawat = :no_rawat AND tgl_perawatan = :tgl_perawatan AND jam_rawat = :jam_rawat
	`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *catatanObservasiRanapRepositoryImpl) Delete(c *fiber.Ctx, noRawat string, tglPerawatan string, jamRawat string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `DELETE FROM catatan_observasi_ranap WHERE no_rawat = $1 AND tgl_perawatan = $2 AND jam_rawat = $3`
	_, err = tx.Exec(query, noRawat, tglPerawatan, jamRawat)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *catatanObservasiRanapRepositoryImpl) FindByNoRawatAndTanggal2(noRawat string, tanggal string) (*entity.CatatanObservasiRanap, error) {
	query := `
        SELECT*
		FROM catatan_observasi_ranap
        WHERE no_rawat = $1 AND tgl_perawatan = $2
        LIMIT 1
    `
	fmt.Println("📦 Executing query for no_rawat =", noRawat, "tgl =", tanggal)
	var result entity.CatatanObservasiRanap
	if err := r.DB.Get(&result, query, noRawat, tanggal); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("catatan observasi tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal query observasi: %v", err)
	}

	return &result, nil
}

func (r *catatanObservasiRanapRepositoryImpl) UpdateByNoRawatAndTanggal(noRawat string, tgl string, e *entity.CatatanObservasiRanap) error {
	query := `
		UPDATE catatan_observasi_ranap
		SET 
			jam_rawat = $1,
			gcs = $2,
			td = $3,
			hr = $4,
			rr = $5,
			suhu = $6
		WHERE no_rawat = $7 AND tgl_perawatan = $8
	`

	_, err := r.DB.Exec(
		query,
		e.JamRawat,
		e.GCS,
		e.TD,
		e.HR,
		e.RR,
		e.Suhu,
		noRawat,
		e.TglPerawatan,
	)
	fmt.Println("🔧 Updating catatan_observasi for", noRawat, tgl)
	fmt.Printf("➡️  Data: %+v\n", e)
	if err != nil {
		return fmt.Errorf("gagal update catatan observasi: %v", err)
	}

	return nil
}
