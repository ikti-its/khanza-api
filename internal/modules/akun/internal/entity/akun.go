package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Akun struct {
	Id        uuid.UUID      `gorm:"column:id;default:uuid_generate_v4()"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	Foto      string         `gorm:"column:foto;default:/assets/img/default.jpg"`
	Role      int            `gorm:"column:role"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater   uuid.UUID      `gorm:"column:updater"`
}

func (Akun) TableName() string {
	return "akun"
}
