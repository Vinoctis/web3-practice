package main 

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-web3-primary-task3/models"
)

func getPostComments(db *gorm.DB, name string) {
	var author models.User 
	db.Debug().Preload("Posts.Comments").Where("name = ?", name).First(&author)
	fmt.Printf("author:%s\n",author.Name)
	for _, post := range author.Posts {
		fmt.Printf("post:%s\n",post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("-- comment:%s\n", comment.Content)
		}
	}
}

func getHotestPost(db *gorm.DB) {
	var post models.Post
	sql := "SELECT * FROM `post` WHERE id = (SELECT post_id FROM `comment` group by post_id order by count(1) desc limit 1)"
	db.Raw(sql).Scan(&post)
	fmt.Printf("The hotest post is :%s\n",post.Title)
}

func AddPost(db *gorm.DB, author *models.User) error {
	post := models.Post{
		Title:"NewPost",
		Content:"NewPostContent",
		UserID: author.ID,
	}
	err := db.Create(&post).Error
	return err
}

func main() {
	host := "localhost"
	user := "vinoctis"
	pwd  := "123456"
	database := "test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err!= nil {
		panic("数据库连接失败")
	}
	// 数据库迁移
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err!= nil {
		panic("数据库迁移失败")
	}

	//获取作者所有的文章和评论
	getPostComments(db, "Vinoctis")

	//获取最热门的文章
	getHotestPost(db)

	//新增文章
	// var author models.User 
	// db.Where("name = ?", "Vinoctis").First(&author)
	// err = AddPost(db, &author)
	// if err!= nil {
	// 	panic("新增文章失败")
	// }

	var post models.Post
	db.Debug().Last(&post)
	var comment models.Comment
	db.Debug().Where("post_id", post.ID).First(&comment)
	db.Debug().Delete(&comment)
}
