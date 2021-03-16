package ucoder

import (
	"crypto/md5"
	"encoding/hex"
)

var Md5 md5Coder

type md5Coder struct{}

func (coder md5Coder) Md5Hex(src string) (des string) {
	h := md5.New()
	h.Write([]byte(src))
	des = hex.EncodeToString(h.Sum(nil))
	return
}
