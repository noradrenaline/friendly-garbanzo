package foo

import "testing"

func TestAndPass(t *testing.T) {
	if 1==0 {
		t.Fail()
	}
}

func TestAndFail(t *testing.T) {
	t.Fail()
}
