package model

import "encoding/json"

type AuthInfo struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}

func (u *AuthInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
