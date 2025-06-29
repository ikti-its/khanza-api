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

type catatanObservasiRanapKebidananRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCatatanObservasiRanapKebidananRepository(db *sqlx.DB) repository.CatatanObservasiRanapKebidananRepository {
	return &catatanObservasiRanapKebidananRepositoryImpl{DB: db}
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userID, ok := userIDRaw.(string)
	if !ok {
		log.Println("⚠️ user_id is not a string")
		return fmt.Errorf("invalid user_id type: expected string, got %T", userIDRaw)
	}

	safeUserID := pq.QuoteLiteral(userID)
	query := fmt.Sprintf(`SET LOCAL my.user_id = %s`, safeUserID)

	if _, err := tx.Exec(query); err != nil {
		log.Printf("❌ Failed to SET LOCAL my.user_id = %v: %v\n", userID, err)
		return err
	}
	return nil
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) Insert(c *fiber.Ctx, data *entity.CatatanObservasiRanapKebidanan) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO catatan_observasi_ranap_kebidanan (
			no_rawat, tgl_perawatan, jam_rawat, gcs, td, hr, rr, suhu, spo2,
			kontraksi, bjj, ppv, vt, nip
		) VALUES (
			:no_rawat, :tgl_perawatan, :jam_rawat, :gcs, :td, :hr, :rr, :suhu, :spo2,
			:kontraksi, :bjj, :ppv, :vt, :nip
		)`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindAll() ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan WHERE no_rawat = $1 ORDER BY tgl_perawatan DESC, jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindByTanggal(tanggal string) ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan WHERE tgl_perawatan = $1 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, tanggal)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) FindByNoRawatAndTanggal(noRawat string, tanggal string) ([]entity.CatatanObservasiRanapKebidanan, error) {
	var list []entity.CatatanObservasiRanapKebidanan
	query := `SELECT * FROM catatan_observasi_ranap_kebidanan WHERE no_rawat = $1 AND tgl_perawatan = $2 ORDER BY jam_rawat DESC`
	err := r.DB.Select(&list, query, noRawat, tanggal)
	return list, err
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) Update(c *fiber.Ctx, data *entity.CatatanObservasiRanapKebidanan) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		UPDATE catatan_observasi_ranap_kebidanan SET
			gcs = :gcs, td = :td, hr = :hr, rr = :rr, suhu = :suhu, spo2 = :spo2,
			kontraksi = :kontraksi, bjj = :bjj, ppv = :ppv, vt = :vt, nip = :nip
		WHERE no_rawat = :no_rawat AND tgl_perawatan = :tgl_perawatan AND jam_rawat = :jam_rawat
	`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *catatanObservasiRanapKebidananRepositoryImpl) Delete(c *fiber.Ctx, noRawat, tglPerawatan, jamRawat string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `DELETE FROM catatan_observasi_ranap_kebidanan WHERE no_rawat = $1 AND tgl_perawatan = $2 AND jam_rawat = $3`
	_, err = tx.Exec(query, noRawat, tglPerawatan, jamRawat)
	if err != nil {
		return err
	}

	return tx.Commit()
}
