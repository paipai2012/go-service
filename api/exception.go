package api

import (
	"encoding/json"
)

type Exception struct {
	Code int `json:"code"`
}

func NewException(code int) string {
	e := Exception{Code: code}
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(data)
}
