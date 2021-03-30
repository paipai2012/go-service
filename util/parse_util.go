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

func ParseConfig(path string) (*model.AppInfo, error) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
