package entity

import "time"

type Role struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Role      string    `gorm:"column:role;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Role) TableName() string {
	return "role"
}
