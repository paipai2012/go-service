package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.GetQuery("token")
		fmt.Println("auth check token..." + token)
		// if !ok {
		// 	panic("token is empty")
		// }
		c.Next()
	}
}
