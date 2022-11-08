package common

import (
	"net/http"
	"sale-service/api"
	"time"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	c.Abort()
}

func FailedParam(c *gin.Context) {
	c.JSON(http.StatusOK, api.JsonError(api.ErrParam))
	c.Abort()
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "Not Found",
		"url":     c.Request.URL.Path,
	})
	c.Abort()
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}
