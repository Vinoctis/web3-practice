package models

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey;type:int(11)"`
	Name string `gorm:"type:varchar(30)"`
	Email string `gorm:"type:varchar(100)"`
	TotalPost int `gorm:"type:int(11)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts []Post
	Comments []Comment
}

// 实现gorm.Tabler接口
func (User) TableName() string {
	return "user"
}