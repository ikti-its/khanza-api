package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Jadwal struct {
	Id        uuid.UUID      `gorm:"column:id;default:uuid_generate_v4()"`
	IdPegawai uuid.UUID      `gorm:"column:id_pegawai"`
	IdHari    int            `gorm:"column:id_hari"`
	IdShift   string         `gorm:"column:id_shift"`
	JamMasuk  string         `gorm:"column:jam_masuk"`
	JamPulang string         `gorm:"column:jam_pulang"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater   uuid.UUID      `gorm:"column:updater"`
}

func (Jadwal) TableName() string {
	return "jadwal_pegawai"
}
