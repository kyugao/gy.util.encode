package ucoder

import (
	"encoding/base64"
)

var Base64 b64

type b64 struct{}

func (coder b64) Encode(src string) (des string) {
	des = base64.StdEncoding.EncodeToString([]byte(src))
	return
}

func (coder b64) Decode(src string) (des string, err error) {
	desByte, err := base64.StdEncoding.DecodeString(src)
	des = string(desByte)
	return
}
