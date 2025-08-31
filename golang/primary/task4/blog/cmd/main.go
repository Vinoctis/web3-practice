package main

import (
	"blog/internal/controller"
	"blog/internal/middleware"
	"blog/internal/model"
	"blog/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	postService := service.NewPostService(db)
	authController := controller.NewAuthController(userService)
	postController := controller.NewPostController(postService)

	//初始化路由
	r := gin.Default()


	public := r.Group("/api/v1")
	{
		//注册
		public.POST("/register", authController.Register)
	}

	//认证路由
	private := r.Group("/api/v1")
	private.Use(middleware.AuthMiddleware())
	{
		//登录
		private.POST("/login", authController.Login)

		//文章相关接口
		post := private.Group("/post")
		{
			post.POST("/create", postController.Create)
			post.POST("/get", postController.Get)
			// post.POST("/update", postController.Update)
			// post.POST("/delete", postController.Delete)
		}
	}
	r.Run(":8080")
}
