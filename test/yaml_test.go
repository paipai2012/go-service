package test

import (
	"io/ioutil"
	"os"
	"sale-service/model"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestTemplate(t *testing.T) {

	type Info struct {
		Name  string
		Age   int
		Hobby []string
	}

	data := `
name: 江景
age: 18
hobby: ["code", "read"]
`

	info := Info{}
	err := yaml.Unmarshal([]byte(data), &info)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(info)
}

func estYaml(t *testing.T) {

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
