package common

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": message,
	})
	c.Abort()
}

func Failed(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}
