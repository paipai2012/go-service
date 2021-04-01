package main

import (
	"context"
	"log"
	"moose-go/engine"
	"moose-go/middleware"
	"moose-go/middleware/recover"
	"moose-go/router"
	"moose-go/util"
	mv "moose-go/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	socketio "github.com/googollee/go-socket.io"
)

var (
	ctx = context.Background()
)

func main() {
	initApp()
}

func initApp() {

	config, err := util.ParseConfig("./config/moose.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	_, err = engine.NewOrmEngine(config)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	rdb := engine.NewRedisHelper()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err.Error())
		return
	}

	gin.SetMode(config.AppMode)

	app := gin.Default()

	// 使用中间件
	app.Use(recover.Recover())
	app.Use(middleware.AuthRequired())

	router.InitRouter(app)

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valuein", mv.ValueIn)
	}

	// init socket io
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	go server.Serve()
	defer server.Close()
	router.InitSocket(app, server)

	app.Run("0.0.0.0:7000")
}
