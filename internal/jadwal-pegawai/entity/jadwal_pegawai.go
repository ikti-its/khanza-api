package entity

import (
	pegawai "github.com/fathoor/simkes-api/internal/pegawai/entity"
	shift "github.com/fathoor/simkes-api/internal/shift/entity"
	"time"
)

type JadwalPegawai struct {
	NIP       string          `gorm:"column:nip;primaryKey"`
	Pegawai   pegawai.Pegawai `gorm:"foreignKey:nip;references:nip"`
	Tahun     int16           `gorm:"column:tahun;primaryKey"`
	Bulan     int16           `gorm:"column:bulan;primaryKey"`
	Hari      int16           `gorm:"column:hari;primaryKey"`
	ShiftNama string          `gorm:"column:shift_nama;not null"`
	Shift     shift.Shift     `gorm:"foreignKey:shift_nama;references:nama"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (JadwalPegawai) TableName() string {
	return "jadwal_pegawai"
}
