package entity

import "time"

type Shift struct {
	Nama      string    `gorm:"column:nama;primaryKey"`
	JamMasuk  time.Time `gorm:"column:jam_masuk;not null"`
	JamKeluar time.Time `gorm:"column:jam_keluar;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Shift) TableName() string {
	return "shift"
}
