package common

import (
	"fmt"
	"moose-go/api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": "请求成功",
	})
}

func Failed(c *gin.Context, e *api.Exception) {
	var err = gin.H{}
	err["code"] = e.Code
	err["message"] = e.Message
	if e.ErrMsg != nil {
		err["errMsg"] = fmt.Sprintf("%v", e.ErrMsg)
	}
	c.JSON(http.StatusOK, err)
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "未授权，请登录！",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "Not Found",
		"url":     c.Request.URL.Path,
	})
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}
