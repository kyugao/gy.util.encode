package uemd5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hex(src string) (des string) {
	h := md5.New()
	h.Write([]byte(src))
	des = hex.EncodeToString(h.Sum(nil))
	return
}
