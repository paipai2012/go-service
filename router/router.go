package router

import (
	v1 "moose-go/controller/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	new(v1.UserController).RegisterRouter(app)
}
