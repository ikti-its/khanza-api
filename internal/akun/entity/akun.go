package entity

import (
	pegawai "github.com/fathoor/simkes-api/internal/pegawai/entity"
	role "github.com/fathoor/simkes-api/internal/role/entity"
	"time"
)

type Akun struct {
	NIP       string          `gorm:"column:nip;primaryKey"`
	Pegawai   pegawai.Pegawai `gorm:"foreignKey:nip;references:nip"`
	Email     string          `gorm:"column:email;unique;not null"`
	Password  string          `gorm:"column:password;not null"`
	RoleNama  string          `gorm:"column:role_nama;not null"`
	Role      role.Role       `gorm:"foreignKey:role_nama;references:nama"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Akun) TableName() string {
	return "akun"
}
