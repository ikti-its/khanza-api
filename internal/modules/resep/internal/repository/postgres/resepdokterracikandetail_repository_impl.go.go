package postgres

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type resepDokterRacikanDetailRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepDokterRacikanDetailRepository(db *sqlx.DB) repository.ResepDokterRacikanDetailRepository {
	return &resepDokterRacikanDetailRepositoryImpl{DB: db}
}

func (r *resepDokterRacikanDetailRepositoryImpl) setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error {
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

func (r *resepDokterRacikanDetailRepositoryImpl) FindAll() ([]entity.ResepDokterRacikanDetail, error) {
	query := `SELECT * FROM resep_dokter_racikan_detail ORDER BY no_resep DESC, no_racik ASC`
	var list []entity.ResepDokterRacikanDetail
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepDokterRacikanDetailRepositoryImpl) FindByNoResepAndNoRacik(noResep, noRacik string) ([]entity.ResepDokterRacikanDetail, error) {
	query := `
		SELECT * FROM resep_dokter_racikan_detail
		WHERE no_resep = $1 AND no_racik = $2
		ORDER BY kode_brng ASC
	`
	var list []entity.ResepDokterRacikanDetail
	err := r.DB.Select(&list, query, noResep, noRacik)
	return list, err
}

func (r *resepDokterRacikanDetailRepositoryImpl) Insert(c *fiber.Ctx, d *entity.ResepDokterRacikanDetail) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO resep_dokter_racikan_detail (
			no_resep, no_racik, kode_brng, p1, p2, kandungan, jml
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`
	_, err = tx.Exec(query,
		d.NoResep, d.NoRacik, d.KodeBrng, d.P1, d.P2, d.Kandungan, d.Jml,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *resepDokterRacikanDetailRepositoryImpl) Update(c *fiber.Ctx, d *entity.ResepDokterRacikanDetail) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		UPDATE resep_dokter_racikan_detail SET 
			p1 = $4,
			p2 = $5,
			kandungan = $6,
			jml = $7
		WHERE no_resep = $1 AND no_racik = $2 AND kode_brng = $3
	`
	_, err = tx.Exec(query,
		d.NoResep, d.NoRacik, d.KodeBrng, d.P1, d.P2, d.Kandungan, d.Jml,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *resepDokterRacikanDetailRepositoryImpl) Delete(c *fiber.Ctx, noResep, noRacik, kodeBrng string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		DELETE FROM resep_dokter_racikan_detail
		WHERE no_resep = $1 AND no_racik = $2 AND kode_brng = $3
	`
	_, err = tx.Exec(query, noResep, noRacik, kodeBrng)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *resepDokterRacikanDetailRepositoryImpl) FindByNoResep(noResep string) ([]model.ResepDokterRacikanDetail, error) {
	var results []model.ResepDokterRacikanDetail

	query := `
		SELECT 
			no_resep, 
			no_racik, 
			kode_brng, 
			p1, 
			p2, 
			kandungan, 
			jml 
		FROM resep_dokter_racikan_detail
		WHERE no_resep = $1
	`

	err := r.DB.Select(&results, query, noResep)
	if err != nil {
		// Hanya log dan return kosong jika tidak ditemukan
		log.Printf("❌ Query failed: %v", err)
		return nil, err
	}

	// Log jika hasilnya kosong tapi tidak error
	log.Printf("✅ Query success. Found %d rows", len(results))
	return results, nil
}
