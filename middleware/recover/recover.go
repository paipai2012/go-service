package recover

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"sale-service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method

				fileName, line, functionName := "?", 0, "?"

				pc, fileName, line, ok := runtime.Caller(2)

				if ok {
					functionName = runtime.FuncForPC(pc).Name()
					functionName = filepath.Ext(functionName)
					functionName = strings.TrimPrefix(functionName, ".")
				}

				log.Printf("异常捕获 | url %s | method | %s | error %v | error type %T | fileName %s | line %d |", url, method, err, err, fileName, line)
				common.JSON(c, fmt.Sprintf("%v", err))
			}
		}()
		c.Next()
	}
}
