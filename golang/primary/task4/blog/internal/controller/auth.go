package controller 

import (
	"github.com/gin-gonic/gin"
	"blog/internal/service"
	"blog/internal/model"
)

type AuthController struct {
	UserService *service.UserService
}

func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{
		UserService: userService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	err := ac.UserService.Register(&model.User{Name: name, Email: email})
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "注册失败",
			"err": err.Error(),
		})
	}
}