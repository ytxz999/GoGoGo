package model

import "time"

type User struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	Username  string
	Password  string
	CreatedAt time.Time
	Todos     []Todo `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
