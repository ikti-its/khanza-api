package entity

import "time"

type Jabatan struct {
	Nama      string    `gorm:"column:nama;primaryKey"`
	Jenjang   string    `gorm:"column:jenjang;not null"`
	GajiPokok float64   `gorm:"column:gaji_pokok;not null;default:0"`
	Tunjangan float64   `gorm:"column:tunjangan;not null;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Jabatan) TableName() string {
	return "jabatan"
}
