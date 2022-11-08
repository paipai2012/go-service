package util

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sale-service/model"

	"gopkg.in/yaml.v2"
)

var _jsonCfg *model.AppInfo = nil

var _yamlCfg *model.AppConfig = nil

func GetJSONConfig() *model.AppInfo {
	return _jsonCfg
}

func ParseJSONConfig(path string) *model.AppInfo {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&_jsonCfg)

	if err != nil {
		panic(err)
	}
	return _jsonCfg
}

func GetYamlConfig() *model.AppConfig {
	return _yamlCfg
}

func ParseYamlConfig(path string) *model.AppConfig {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	err = yaml.Unmarshal([]byte(data), &_yamlCfg)

	if err != nil {
		panic(err)
	}

	log.Println(_yamlCfg)

	return _yamlCfg
}
