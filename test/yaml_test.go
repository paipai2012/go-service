package test

import (
	"io/ioutil"
	"moose-go/model"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestYaml(t *testing.T) {

	file, err := os.Open("../config/application-dev.yml")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err != nil {
		t.Log(err)
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		t.Log(err)
		return
	}

	t.Log(string(data))

	appConfig := model.AppConfig{}
	yaml.Unmarshal([]byte(data), &appConfig)
	t.Log(appConfig.MySql)
	t.Log(appConfig.Redis)

}
