package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Pegawai struct {
	Id           uuid.UUID      `gorm:"column:id;default:uuid_generate_v4()"`
	IdAkun       uuid.UUID      `gorm:"column:id_akun"`
	NIP          string         `gorm:"column:nip"`
	Nama         string         `gorm:"column:nama"`
	JenisKelamin string         `gorm:"column:jenis_kelamin"`
	Jabatan      int            `gorm:"column:id_jabatan"`
	Departemen   int            `gorm:"column:id_departemen"`
	StatusAktif  string         `gorm:"column:id_status_aktif"`
	JenisPegawai string         `gorm:"column:jenis_pegawai"`
	Telepon      string         `gorm:"column:telepon"`
	TanggalMasuk time.Time      `gorm:"column:tanggal_masuk"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater      uuid.UUID      `gorm:"column:updater"`
}

func (Pegawai) TableName() string {
	return "pegawai"
}
