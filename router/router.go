package router

import (
	"sale-service/common"
	"sale-service/controller"

	// v1 "sale-service/controller/v1"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func InitRouter(app *gin.Engine) {
	new(controller.LuckController).RegisterRouter(app)
	new(controller.AccountController).RegisterRouter(app)
	// new(controller.ChatgptController).RegisterRouter(app)
	controller.ChatgptController.RegisterRouter(app)

	// new(v1.UserController).RegisterRouter(app)
	// new(controller.QRCodeController).RegisterRouter(app)
	// new(controller.SmsController).RegisterRouter(app)
	// new(controller.SignInController).RegisterRouter(app)

	app.NoRoute(common.NotFound)
}

func InitSocket(app *gin.Engine, socket *socketio.Server) {
	// new(controller.SocketController).RegisterSocket(app, socket)
}
