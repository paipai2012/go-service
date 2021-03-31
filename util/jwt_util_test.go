package util

import "testing"

// 普通测试
func TestCreateJWT(t *testing.T) {
	jwtStr, err := CreateJWT()
	if err != nil {
		t.Error()
	}
	t.Log(jwtStr)
}
