package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Alamat struct {
	IdAkun    uuid.UUID      `gorm:"column:id_akun"`
	Alamat    string         `gorm:"column:alamat"`
	AlamatLat float64        `gorm:"column:alamat_lat"`
	AlamatLon float64        `gorm:"column:alamat_lon"`
	Kota      string         `gorm:"column:kota"`
	KodePos   string         `gorm:"column:kode_pos"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater   uuid.UUID      `gorm:"column:updater"`
}

func (Alamat) TableName() string {
	return "alamat"
}
