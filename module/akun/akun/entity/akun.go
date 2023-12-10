package entity

import (
	"github.com/fathoor/simkes-api/module/akun/role/entity"
	"time"
)

type Akun struct {
	NIP       string      `gorm:"column:nip;primaryKey"`
	Email     string      `gorm:"column:email;unique;not null"`
	Password  string      `gorm:"column:password;not null"`
	RoleName  string      `gorm:"column:role_name;not null"`
	Role      entity.Role `gorm:"foreignKey:role_name;references:name"`
	CreatedAt time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Akun) TableName() string {
	return "akun"
}
