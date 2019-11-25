package verify

import "testing"

func TestVerfifyMobile(t *testing.T) {
	m1 := "15313511069"
	m2 := "3511069"
	m3 := "153135110691"
	m4 := "351106911111111"

	if VerfifyMobile(m1) {
		t.Log()
	} else {
		t.Fail()
	}

	if !VerfifyMobile(m2) {
		t.Log()
	} else {
		t.Fail()
	}

	if !VerfifyMobile(m3) {
		t.Log()
	} else {
		t.Fail()
	}

	if !VerfifyMobile(m4) {
		t.Log()
	} else {
		t.Fail()
	}
}
