package controller

import (
	"blog/internal/model"
	"blog/internal/response"
	"blog/internal/service"
	"blog/internal/utils"
	"fmt"

	"github.com/gin-gonic/gin"
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
		response.Fail(c, "注册失败:"+err.Error(), nil)
	} else {
		response.Ok(c, nil)
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq struct {
		Email string `json:"email" binding:"required,email"` 
	}

	if err := c.ShouldBindJSON(&loginReq);err != nil {
		response.Fail(c, "登录失败,请求参数有误:"+err.Error(), nil)
		return
	}
	user, err  := ac.UserService.Login(loginReq.Email)
	if err != nil {
		response.Fail(c, "登录失败，用户验证失败", nil)
		return
	} 
	//生成JWT token
	token, err := utils.GenerateToken(int(user.ID), user.Name)
	if err != nil {
		response.Fail(c, "登录失败，token生成失败"+err.Error(), nil)
		return
	}
	claims, _ := c.Get("claims")
	fmt.Printf("claims: %v", claims)
	//返回token
	response.Ok(c, gin.H{
		"token": token, 
		"user": gin.H{"id": user.ID, "name": user.Name, "email": user.Email},
	})
	return
}
