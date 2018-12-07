package dog_test

import (
	"fmt"
	"github.com/progayk/gowebdev/go-programming/exercises-level_12/01/dog"
	"testing"
)

func TestConvToDogYears(t *testing.T) {
	var v float64
	v = dog.ConvToDogYears(35)
	if v != float64(5) {
		t.Errorf("expected 5 got %v", v)
	}
}

func ExampleConvToDogYears() {
	humanYear := 35
	fmt.Println(dog.ConvToDogYears(humanYear))
	// Output:
	// 5
}
