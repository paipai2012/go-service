package app

import (
	"log"
	"os"
	"os/signal"
	"sale-service/engine"
	"sale-service/middleware/recover"
	"sale-service/router"
	"sale-service/util"
	"syscall"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func InitGin() {
	config := util.GetYamlConfig()
	app := gin.Default()

	// 使用中间件
	app.Use(recover.Recover())
	// app.Use(middleware.AuthRequired())

	router.InitRouter(app)

	// init socket io
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	go server.Serve()
	defer server.Close()
	router.InitSocket(app, server)

	gin.SetMode(config.App.Mode)

	handleSignal(app)

	app.Run("0.0.0.0:7000")
}

func handleSignal(server *gin.Engine) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Printf("got signal [%s], exiting now", s)

		engine.CloseOrmEngine()

		os.Exit(0)
	}()
}
