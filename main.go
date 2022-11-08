package main

import (
	"sale-service/app"
	"sale-service/util"

	"github.com/gin-gonic/gin"
)

func init() {
	mode := gin.Mode()
	if mode == "release" {
		util.ParseYamlConfig("./config/application-prod.yml")
	} else {
		util.ParseYamlConfig("./config/application-dev.yml")
	}
}

func main() {

	app.InitEngine()

	app.BindValidator()

	app.InitGin()
}
