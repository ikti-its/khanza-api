package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Foto struct {
	IdPegawai uuid.UUID      `gorm:"column:id_pegawai"`
	Foto      string         `gorm:"column:foto"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater   uuid.UUID      `gorm:"column:updater"`
}

func (Foto) TableName() string {
	return "foto_pegawai"
}
