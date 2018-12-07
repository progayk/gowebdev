package main

import "testing"

func TestMyAdd(t *testing.T) {
	v := myAdd(1,2)
	if v != 3 {
		t.Error("Expected value is 3 got", v)
	}
}

