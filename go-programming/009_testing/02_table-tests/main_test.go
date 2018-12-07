package main

import "testing"


func TestMyAdd(t *testing.T) {
	type testPair struct {
		Args []int
		Result int
	}

	tests := []testPair{
		{Args: []int{2, 3}, Result: 5},
		{Args: []int{200, 301}, Result: 501},
		{Args: []int{-2, 3}, Result: 1},
	}

	for _, v := range tests {
		r := myAdd(v.Args...)
		if r != v.Result {
			t.Error("Expected value is 5 got", r)
		}
	}
}

