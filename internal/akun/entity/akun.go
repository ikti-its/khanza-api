package entity

import (
	"github.com/fathoor/simkes-api/internal/role/entity"
	"time"
)

type Akun struct {
	NIP       string      `gorm:"column:nip;primaryKey"`
	Email     string      `gorm:"column:email;unique;not null"`
	Password  string      `gorm:"column:password;not null"`
	RoleNama  string      `gorm:"column:role_nama;not null"`
	Role      entity.Role `gorm:"foreignKey:role_nama;references:nama"`
	CreatedAt time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Akun) TableName() string {
	return "akun"
}
