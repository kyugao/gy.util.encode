package uebase64

import (
	"encoding/base64"
)

func Encode(src string) (des string){
	des = base64.StdEncoding.EncodeToString([]byte(src))
	return
}

func Decode(src string) (des string, err error) {
	desByte, err := base64.StdEncoding.DecodeString(src)
	des = string(desByte)
	return
}