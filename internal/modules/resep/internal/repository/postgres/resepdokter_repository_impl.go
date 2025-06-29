package postgres

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type resepDokterRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepDokterRepository(db *sqlx.DB) repository.ResepDokterRepository {
	return &resepDokterRepositoryImpl{DB: db}
}

func (r *resepDokterRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
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

func (r *resepDokterRepositoryImpl) FindAll() ([]entity.ResepDokter, error) {
	query := `SELECT * FROM resep_dokter ORDER BY no_resep DESC`
	var list []entity.ResepDokter
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepDokterRepositoryImpl) FindByNoResep(noResep string) ([]entity.ResepDokter, error) {
	query := `SELECT * FROM resep_dokter WHERE no_resep = $1 ORDER BY kode_barang ASC`
	var list []entity.ResepDokter
	err := r.DB.Select(&list, query, noResep)
	return list, err
}

func (r *resepDokterRepositoryImpl) Insert(c *fiber.Ctx, p *entity.ResepDokter) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO resep_dokter (
			no_resep, kode_barang, jumlah, aturan_pakai
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err = tx.Exec(query, p.NoResep, p.KodeBarang, p.Jumlah, p.AturanPakai)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *resepDokterRepositoryImpl) Update(c *fiber.Ctx, p *entity.ResepDokter) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		UPDATE resep_dokter SET 
			jumlah = $3, aturan_pakai = $4
		WHERE no_resep = $1 AND kode_barang = $2
	`
	_, err = tx.Exec(query, p.NoResep, p.KodeBarang, p.Jumlah, p.AturanPakai)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *resepDokterRepositoryImpl) Delete(c *fiber.Ctx, noResep, kodeBarang string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `DELETE FROM resep_dokter WHERE no_resep = $1 AND kode_barang = $2`
	_, err = tx.Exec(query, noResep, kodeBarang)
	if err != nil {
		return err
	}

	return tx.Commit()
}
