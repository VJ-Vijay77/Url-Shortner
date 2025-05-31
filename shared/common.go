package shared

import (
	"encoding/base64"
)

func Encodetobase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Decodetobase64(s string) (string, error) {
	ds, err := base64.StdEncoding.DecodeString(s)
	return string(ds), err
}
