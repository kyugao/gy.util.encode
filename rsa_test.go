package ucoder

import (
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	pri, pub, _ := MakeSSHKeyPair()
	fmt.Println(pri)
	t.Log(pub)
}
