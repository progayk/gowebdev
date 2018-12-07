package acdc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	v := Sum(1,2)
	if v != 3 {
		t.Error("expected 3 got", v)
	}
}

func ExampleSum() {
	fmt.Println(Sum(1,2))
	// Output:
	// 3
}


