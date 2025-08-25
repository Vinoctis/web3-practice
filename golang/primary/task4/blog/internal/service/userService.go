package service

import (
	"blog/internal/model"
	"gorm.io/gorm"
	"errors"
)

type UserService struct{
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
} 

func (us *UserService) Register(user *model.User) error {
	//
	var ExistUser model.User
	us.DB.Where("name = ? or email = ?", user.Name, user.Email).First(&ExistUser)
	if ExistUser.ID != 0 {
		return errors.New("用户已存在")
	}
	if err := us.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) Login(email string) (*model.User, error) {
	var user model.User
	err := us.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
	}
	return &user, nil
}


