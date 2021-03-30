package middleware

import (
	"encoding/json"
	"log"
	"moose-go/api"
	"moose-go/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method
				log.Printf("| url [%s] | method | [%s] | error [%s] |", url, method, err)
				var exception api.Exception
				err := json.Unmarshal([]byte(string(err.(string))), &exception)
				if err != nil {
					common.Failed(c, http.StatusBadRequest, "未知错误，请联系管理员！")
					c.Abort()
					return
				}
				// 没有定义
				errorMessage, ok := api.StatusText(exception.Code)
				if !ok {
					errorMessage = "系统异常"
				}
				common.Failed(c, exception.Code, errorMessage)
				c.Abort()
			}
		}()
		c.Next()
	}
}
