package main

import (
	"moose-go/app"
	"moose-go/util"
)

func init() {
	util.ParseConfig("./config/moose.json")
}

func main() {

	app.InitEngine()

	app.BindValidator()

	app.InitGin()
}
