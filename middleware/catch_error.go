package middleware

import (
	"encoding/json"
	"fmt"
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
				log.Printf("异常捕获 | url %s | method | %s | error %s |", url, method, err)

				if err != nil {
					switch err.(type) {
					case string:
						var exception api.Exception
						err := json.Unmarshal([]byte(string(err.(string))), &exception)

						// parse err; not api.Exception
						if err != nil {
							common.Failed(c, http.StatusBadRequest, err.Error())
							return
						}
						// 没有定义
						errorMessage, ok := api.StatusText(exception.Code)
						if !ok {
							errorMessage = "未知错误！"
						}
						common.Failed(c, exception.Code, errorMessage)
					default:
						common.Failed(c, -1, fmt.Sprintf("%s %v", "系统异常", err))
					}
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
