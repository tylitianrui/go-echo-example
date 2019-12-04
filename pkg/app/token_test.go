package app

import "testing"

func TestGenToken(t *testing.T) {
	if s, err := GenToken(1); err != nil {
		t.Fatal(err)
	} else {
		t.Log(s)
	}

}
