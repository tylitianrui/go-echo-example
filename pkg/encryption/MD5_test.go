package encryption

import (
	"fmt"
	"testing"
)

func TestMD5Encrypt_Compare(t *testing.T) {
	md := NewMD5()
	c := md.Encode("12345")
	fmt.Println(c)
	//	3132333435d41d8cd98f00b204e9800998ecf8427e
	if !md.Compare("12345", "3132333435d41d8cd98f00b204e9800998ecf8427e") {
		t.Fail()
	}

}
