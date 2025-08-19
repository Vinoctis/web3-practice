package service

import (
	"blog/internal/model"
	"gorm.io/gorm"
)

type UserService struct{
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
} 

func (us *UserService) Register(user *model.User) error {
	us.DB.Create(user)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// func (us *UserService) Login(name string) error {
// 	var user model.User
// 	us.DB.Where("username = ? and password = ?", name).First(&user)
// 	return nil
// }


