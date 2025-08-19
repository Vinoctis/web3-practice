package model

import (
	"time"
	"gorm.io/gorm"
)

type Post struct {
	ID uint `gorm:"primaryKey;type:int(11)"`
	Title string `gorm:"type:varchar(50)"`
	Content string `gorm:"type:text"`
	UserID uint `gorm:"type:int(11)"`
	Status string `gorm:"type:varchar(50)"`
	User User
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Comments []Comment
}

func (Post) TableName() string {
	return "post"
}

func (p Post) AfterCreate(tx *gorm.DB) (err error){
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("total_post", gorm.Expr("total_post + 1"))
	return 
}
