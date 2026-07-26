package model

import "time"

// 待办事项状态
const (
	TodoStatusPending = 0
	TodoStatusDone    = 1
)

type Todo struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Title     string
	Content   string
	Status    int
	CreatedAt time.Time
	StartTime time.Time
	EndTime   time.Time
}

func (Todo) TableName() string {
	return "todos"
}
