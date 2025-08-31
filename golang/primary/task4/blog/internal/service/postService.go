package service

import (
	"blog/internal/model"
	//"context"
	"errors"
	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		DB: db,
	}
}

type PostQuery struct {
	ID *uint 
	Title *string
}

func (ps *PostService) Post(post *model.Post) error {
	result := gorm.WithResult()
	err := gorm.G[model.Post](ps.DB, result).Create(nil, post)
	if err != nil {
		return err 
	}
	return nil
}

func (ps *PostService) Get(q *PostQuery) (*model.Post, error) {
	//ctx := context.Background()
	var post model.Post
	query := ps.DB.Model(&post).Debug()
	if q.ID != nil {
		query = query.Where("id = ?", *q.ID)
	}
	if q.Title != nil {
		query = query.Where("title like ?", "%" + *q.Title + "%")
	}
	err := query.First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("查找的文章不存在")
		}
		return nil,err 
	}
	return &post, nil
}