package util

import (
	"crypto/sha1"

	"encoding/hex"
)

func CryptoSha1(s string) string {
	s1 := sha1.New()
	s1.Write([]byte(s))
	return hex.EncodeToString(s1.Sum(nil))
}
