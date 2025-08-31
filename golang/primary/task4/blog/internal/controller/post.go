package controller

import (
	"blog/internal/response"
	"blog/internal/service"
	"blog/internal/model"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

func (pc *PostController) Create(c *gin.Context) {
	var postReq struct {
		Title string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&postReq);err!=nil {
		response.Fail(c,"发布失败：" + err.Error(), nil)
		return
	}
	var post  = model.Post{
		Title: postReq.Title, Content: postReq.Content,
	}
	err := pc.PostService.Post(&post)
	if err != nil {
		response.Fail(c, "发布失败："+err.Error(), nil)
		return 
	}
	response.Ok(c, gin.H{
		"post_id": post.ID,
		"post_title": post.Title,
	})
}

func (pc *PostController) Get(c *gin.Context) {
	var postReq struct {
		ID uint `json:"id"`
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&postReq);err!=nil {
		response.Fail(c,"查询败：" + err.Error(), nil)
		return
	}
	var query service.PostQuery
    if postReq.ID != 0 {  // 只在非零时设置
        query.ID = &postReq.ID
    }
    if postReq.Title != "" {
        query.Title = &postReq.Title
    }
	post, err := pc.PostService.Get(&query)
	if err != nil {
		response.Fail(c,err.Error(), nil)
		return 
	}
	response.Ok(c, gin.H{
		"post_id": post.ID,
		"post_title": post.Title,
	})
}