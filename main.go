package main

import (
	"moose-go/app"
	"moose-go/util"
)

func init() {
	// util.ParseJSONConfig("./config/application-dev.json")
	util.ParseYamlConfig("./config/application-dev.yml")
	// util.ParseYamlConfig("./config/application-prod.yml")
}

func main() {

	app.InitEngine()

	app.BindValidator()

	app.InitGin()
}
