package models

import (
	"time"
	"gorm.io/gorm"
)

type Comment struct {
	ID uint `gorm:"primaryKey;type:int(11)"`
	Content string `gorm:"type:text"`
	UserID uint `gorm:"type:int(11)"`
	User User
	PostID uint `gorm:"type:int(11)"`
	//Post Post `gorm:"references:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Comment) TableName() string {
	return "comment"
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("status", "无评论")
	}
	return
}

