package entity

import "time"

type Role struct {
	Nama      string    `gorm:"column:nama;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Role) TableName() string {
	return "role"
}
