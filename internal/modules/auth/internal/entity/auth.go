package entity

import "github.com/google/uuid"

type Auth struct {
	Id       uuid.UUID `gorm:"column:id"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`
	Role     int       `gorm:"column:role"`
}

func (Auth) TableName() string {
	return "akun"
}
