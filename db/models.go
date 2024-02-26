// db/models.go
package db

import "time"

// User 用户模型
type User struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Todo 待办事项模型
type Todo struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	Content   string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
