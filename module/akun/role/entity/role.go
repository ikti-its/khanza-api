package entity

type Role struct {
	ID   int    `gorm:"column:id;primaryKey;autoIncrement"`
	Role string `gorm:"column:role;not null"`
}

func (Role) TableName() string {
	return "role"
}
