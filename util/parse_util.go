package util

import (
	"bufio"
	"encoding/json"
	"moose-go/model"
	"os"
)

var _cfg *model.AppInfo = nil

func GetConfig() *model.AppInfo {
	return _cfg
}

func ParseConfig(path string) *model.AppInfo {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&_cfg)

	if err != nil {
		panic(err)
	}
	return _cfg
}
