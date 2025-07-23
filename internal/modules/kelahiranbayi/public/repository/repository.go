package repository

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/kelahiranbayi/public/entity"
)

type Repository interface {
	FindAll() ([]entity.Entity, error)
	FindById(id string) (entity.Entity, error)
	Insert(c *fiber.Ctx, entity *entity.Entity) error
	Update(c *fiber.Ctx, entity *entity.Entity) error
	Delete(c *fiber.Ctx, id string) error 
    setUserAuditContext(tx *sqlx.Tx, c *fiber.Ctx) error
	UpdateIfExists(c *fiber.Ctx, entity *entity.Entity) error
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

	ip_address_Raw := c.Locals("ip_address")
	ip_address, ok2 := ip_address_Raw.(string)
	if !ok2 {
		return fmt.Errorf("invalid ip_address type: %T", ip_address_Raw)
	}
	safe_ip_address := pq.QuoteLiteral(ip_address)
	_, err = tx.Exec(fmt.Sprintf(`SET LOCAL my.ip_address = %s`, safe_ip_address))

	mac_address_Raw := c.Locals("mac_address")
	mac_address, ok2 := mac_address_Raw.(string)
	if !ok2 {
		return fmt.Errorf("invalid mac_address type: %T", mac_address_Raw)
	}
	safe_mac_address := pq.QuoteLiteral(mac_address)
	_, err = tx.Exec(fmt.Sprintf(`SET LOCAL my.mac_address = %s`, safe_mac_address))


	_, err = tx.Exec(fmt.Sprintf(`SET LOCAL my.encryption_key = %s`, c.Locals("encryption_key")))
	
	return err
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.Entity, error) {
	query := `
		SELECT * FROM kelahiran_bayi ORDER BY no_rkm_medis DESC
	`
	var records []entity.Entity
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(id string) (entity.Entity, error) {
	query := `SELECT * FROM kelahiran_bayi WHERE no_rkm_medis = $1`

	var record entity.Entity
	err := r.DB.Get(&record, query, id)
	return record, err
}

func (r *RepositoryImpl) Insert(c *fiber.Ctx, entity *entity.Entity) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	query := `
		INSERT INTO kelahiran_bayi (
			no_rkm_medis, nm_pasien, jk, tmp_lahir, tgl_lahir, jam, umur, tgl_daftar,
			nm_ibu, umur_ibu, nm_ayah, umur_ayah, alamat,
			bb, pb, proses_lahir, kelahiran_ke, keterangan, diagnosa, penyulit_kehamilan, ketuban,
			lk_perut, lk_kepala, lk_dada, penolong, no_skl, gravida, para, abortus,
			f1, u1, t1, r1, w1, n1,
			f5, u5, t5, r5, w5, n5,
			f10, u10, t10, r10, w10, n10,
			resusitas, obat, mikasi, mikonium
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11, $12,
			$13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27,
			$28, $29, $30, $31, $32, $33,
			$34, $35, $36, $37, $38, $39,
			$40, $41, $42, $43, $44, $45,
			$46, $47, $48, $49, $50, $51
		)
	`

	_, err = tx.Exec(query,
		entity.No_rkm_medis, entity.Nm_pasien, entity.Jk, entity.Tmp_lahir, entity.Tgl_lahir, entity.Jam, entity.Umur, entity.Tgl_daftar,
		entity.Nm_ibu, entity.Umur_ibu, entity.Nm_ayah, entity.Umur_ayah, entity.Alamat,
		entity.Bb, entity.Pb, entity.Proses_lahir, entity.Kelahiran_ke, entity.Keterangan, entity.Diagnosa, entity.Penyulit_kehamilan, entity.Ketuban,
		entity.Lk_perut, entity.Lk_kepala, entity.Lk_dada, entity.Penolong, entity.No_skl, entity.Gravida, entity.Para, entity.Abortus,
		entity.F1, entity.U1, entity.T1, entity.R1, entity.W1, entity.N1,
		entity.F5, entity.U5, entity.T5, entity.R5, entity.W5, entity.N5,
		entity.F10, entity.U10, entity.T10, entity.R10, entity.W10, entity.N10,
		entity.Resusitas, entity.Obat, entity.Mikasi, entity.Mikonium,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *RepositoryImpl) Update(c *fiber.Ctx, entity *entity.Entity) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}
	query := `
		UPDATE kelahiran_bayi SET
			nm_pasien = $2, jk = $3, tmp_lahir = $4, tgl_lahir = $5, jam = $6, umur = $7, tgl_daftar = $8,
			nm_ibu = $9, umur_ibu = $10, nm_ayah = $11, umur_ayah = $12, alamat = $13,
			bb = $14, pb = $15, proses_lahir = $16, kelahiran_ke = $17, keterangan = $18, diagnosa = $19,
			penyulit_kehamilan = $20, ketuban = $21, lk_perut = $22, lk_kepala = $23, lk_dada = $24, penolong = $25, no_skl = $26,
			gravida = $27, para = $28, abortus = $29,
			f1 = $30, u1 = $31, t1 = $32, r1 = $33, w1 = $34, n1 = $35,
			f5 = $36, u5 = $37, t5 = $38, r5 = $39, w5 = $40, n5 = $41,
			f10 = $42, u10 = $43, t10 = $44, r10 = $45, w10 = $46, n10 = $47,
			resusitas = $48, obat = $49, mikasi = $50, mikonium = $51
		WHERE no_rkm_medis = $1
	`
	_, err = tx.Exec(query,
		entity.No_rkm_medis, entity.Nm_pasien, entity.Jk, entity.Tmp_lahir, entity.Tgl_lahir, entity.Jam, entity.Umur, entity.Tgl_daftar,
		entity.Nm_ibu, entity.Umur_ibu, entity.Nm_ayah, entity.Umur_ayah, entity.Alamat,
		entity.Bb, entity.Pb, entity.Proses_lahir, entity.Kelahiran_ke, entity.Keterangan, entity.Diagnosa, entity.Penyulit_kehamilan, entity.Ketuban,
		entity.Lk_perut, entity.Lk_kepala, entity.Lk_dada, entity.Penolong, entity.No_skl, entity.Gravida, entity.Para, entity.Abortus,
		entity.F1, entity.U1, entity.T1, entity.R1, entity.W1, entity.N1,
		entity.F5, entity.U5, entity.T5, entity.R5, entity.W5, entity.N5,
		entity.F10, entity.U10, entity.T10, entity.R10, entity.W10, entity.N10,
		entity.Resusitas, entity.Obat, entity.Mikasi, entity.Mikonium,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *RepositoryImpl) UpdateIfExists(c *fiber.Ctx, bayi *entity.Entity) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Audit context (jika ada)
	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

	// Cek apakah record ada
	var count int
	err = tx.Get(&count, "SELECT COUNT(*) FROM kelahiran_bayi WHERE no_rkm_medis = $1", bayi.No_rkm_medis)
	if err != nil {
		return err
	}
	if count == 0 {
		return nil // Skip kalau data tidak ditemukan
	}

	// Update data jika ditemukan
	query := `
		UPDATE kelahiran_bayi SET 
			nm_pasien = $1,
			jk = $2,
			tmp_lahir = $3,
			tgl_lahir = $4,
			alamat = $5,
			nm_ibu = $6
		WHERE no_rkm_medis = $7
	`
	_, err = tx.Exec(query,
		bayi.Nm_pasien,
		bayi.Jk,
		bayi.Tmp_lahir,
		bayi.Tgl_lahir,
		bayi.Alamat,
		bayi.Nm_ibu,
		bayi.No_rkm_medis,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}


func (r *RepositoryImpl) Delete(c *fiber.Ctx, id string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.setUserAuditContext(tx, c); err != nil {
		return err
	}

    query := `
		DELETE FROM kelahiran_bayi WHERE no_rkm_medis = $1
	`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
