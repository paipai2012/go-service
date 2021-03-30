package middleware

import (
	"log"
	"moose-go/common"

	"github.com/gin-gonic/gin"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method
				log.Printf("| url [%s] | method | [%s] | error [%s] |", url, method, err)

				common.Failed(c, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}
