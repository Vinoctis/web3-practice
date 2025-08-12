package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data any `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func Result(ctx *gin.Context,code int, msg string, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(ctx *gin.Context, data any) {
	Result(ctx, SUCCESS, "success", data)
}

func Fail(ctx *gin.Context, msg string, data any) {
	Result(ctx, ERROR, msg, data)
}




