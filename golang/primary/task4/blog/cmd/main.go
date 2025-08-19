package main 

import (
	"fmt"
	"blog/internal/service"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"blog/internal/model"
	"blog/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	host := "127.0.0.1"
	username := "root"
	password := "123456"
	port := "33061"
	dbname := "blog"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	//初始化服务
	userService := service.NewUserService(db)
	authController := controller.NewAuthController(userService)

	//初始化路由
	r := gin.Default()
	
	public := r.Group("/api/v1")
	{
		public.POST("/register", authController.Register)
	}

	r.Run(":8080")
}
