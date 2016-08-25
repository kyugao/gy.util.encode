package encode

import (
	"testing"
	"util/encode/base64"
	"util/encode/md5"
)

func TestBase64(t *testing.T) {
	srcPwd := "GW@2016"
	b64Pwd := base64.Encode(srcPwd)
	t.Logf("src password  : %s", srcPwd)
	t.Logf("base64 encoded: %s", b64Pwd)


	decPwd, err := base64.Decode(b64Pwd)
	if err != nil {
		t.Logf("decode base64 error %s.", err.Error())
	}
	md5Pwd := md5.Md5Hex(decPwd)
	t.Logf("src password  : %s", decPwd)
	t.Logf("md5 encoded: %s", md5Pwd)
}