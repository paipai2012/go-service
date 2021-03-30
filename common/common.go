package common

import (
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

func Failed(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "未授权，请登录！",
	})
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}
