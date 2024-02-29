package entity

import (
	pegawai "github.com/fathoor/simkes-api/internal/pegawai/entity"
	shift "github.com/fathoor/simkes-api/internal/shift/entity"
	"github.com/google/uuid"
	"time"
)

type Kehadiran struct {
	ID         uuid.UUID       `gorm:"column:id;primaryKey"`
	NIP        string          `gorm:"column:nip;not null"`
	Pegawai    pegawai.Pegawai `gorm:"foreignKey:nip;references:nip"`
	Tanggal    time.Time       `gorm:"column:tanggal;not null"`
	ShiftNama  string          `gorm:"column:shift_nama;not null"`
	Shift      shift.Shift     `gorm:"foreignKey:shift_nama;references:nama"`
	JamMasuk   time.Time       `gorm:"column:jam_masuk;not null"`
	JamKeluar  time.Time       `gorm:"column:jam_keluar;not null"`
	Keterangan string          `gorm:"column:keterangan"`
	CreatedAt  time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Kehadiran) TableName() string {
	return "kehadiran"
}
