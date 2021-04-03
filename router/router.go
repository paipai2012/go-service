package router

import (
	"moose-go/controller"
	v1 "moose-go/controller/v1"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func InitRouter(app *gin.Engine) {
	new(v1.UserController).RegisterRouter(app)
	new(controller.AccountController).RegisterRouter(app)
	new(controller.QRCodeController).RegisterRouter(app)
	new(controller.SmsController).RegisterRouter(app)
}

func InitSocket(app *gin.Engine, socket *socketio.Server) {
	// new(controller.SocketController).RegisterSocket(app, socket)
}
