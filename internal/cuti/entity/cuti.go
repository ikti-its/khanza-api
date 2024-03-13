package entity

import (
	"github.com/fathoor/simkes-api/internal/pegawai/entity"
	"github.com/google/uuid"
	"time"
)

type Cuti struct {
	ID             uuid.UUID      `gorm:"column:id;primaryKey"`
	NIP            string         `gorm:"column:nip;not null"`
	Pegawai        entity.Pegawai `gorm:"foreignKey:nip;references:nip"`
	TanggalMulai   time.Time      `gorm:"column:tanggal_mulai;not null"`
	TanggalSelesai time.Time      `gorm:"column:tanggal_selesai;not null"`
	Keterangan     string         `gorm:"column:keterangan;not null"`
	Status         bool           `gorm:"column:status;not null;default:false"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Cuti) TableName() string {
	return "cuti"
}
