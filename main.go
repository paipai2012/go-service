package main

import (
	"fmt"
	"log"
	"moose-go/engine"
	"moose-go/router"
	"moose-go/util"

	"github.com/gin-gonic/gin"
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

	_, err = engine.OrmEngine(config)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	gin.SetMode(config.AppMode)

	app := gin.Default()

	// app.Use(Auth())

	router.InitRouter(app)

	app.Run("0.0.0.0:8090")
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("auth ...")
	}
}
