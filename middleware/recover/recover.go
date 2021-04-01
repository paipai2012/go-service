package recover

import (
	"fmt"
	"log"
	"moose-go/api"
	"moose-go/common"
	"path/filepath"
	"runtime"
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

				if err != nil {
					switch err.(type) {
					case *api.Exception:
						common.Failed(c, err.(*api.Exception))
					case string:
						common.Failed(c, &api.Exception{Code: -1, Message: err.(string)})
					default:
						common.Failed(c, &api.Exception{Code: -2, Message: fmt.Sprintf("%v", err)})
					}
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
