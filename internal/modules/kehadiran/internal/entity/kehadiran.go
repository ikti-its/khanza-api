package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Kehadiran struct {
	Id              uuid.UUID      `gorm:"column:id;default:uuid_generate_v4()"`
	IdPegawai       uuid.UUID      `gorm:"column:id_pegawai"`
	IdJadwalPegawai uuid.UUID      `gorm:"column:id_jadwal_pegawai"`
	Tanggal         time.Time      `gorm:"column:tanggal"`
	JamMasuk        time.Time      `gorm:"column:jam_masuk"`
	JamPulang       time.Time      `gorm:"column:jam_pulang"`
	Keterangan      string         `gorm:"column:keterangan"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater         uuid.UUID      `gorm:"column:updater"`
}

func (Kehadiran) TableName() string {
	return "presensi"
}
