package router

import (
	"moose-go/common"
	"moose-go/controller"

	// v1 "moose-go/controller/v1"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func InitRouter(app *gin.Engine) {
	new(controller.AccountController).RegisterRouter(app)
	new(controller.ArticleController).RegisterRouter(app)

	// new(v1.UserController).RegisterRouter(app)
	// new(controller.QRCodeController).RegisterRouter(app)
	// new(controller.SmsController).RegisterRouter(app)
	// new(controller.SignInController).RegisterRouter(app)

	app.NoRoute(common.NotFound)
}

func InitSocket(app *gin.Engine, socket *socketio.Server) {
	// new(controller.SocketController).RegisterSocket(app, socket)
}
