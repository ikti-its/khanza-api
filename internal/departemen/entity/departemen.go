package entity

import "time"

type Departemen struct {
	Nama      string    `gorm:"column:nama;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Departemen) TableName() string {
	return "departemen"
}
