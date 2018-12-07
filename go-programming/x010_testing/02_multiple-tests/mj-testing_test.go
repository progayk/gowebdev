package mjmath

import (
	"fmt"
	"testing"
)

type testpair struct {
	values []float64
	avarage float64
}

var tests = []testpair{
	{[]float64{1,2}, 1.5},
	{[]float64{1,1,1,1,1}, 1},
	{[]float64{-1,1}, 0},
}

func TestAvarage(t *testing.T) {
	for _, pair := range tests {
		v := Avarage(pair.values)
		if v != pair.avarage {
			t.Error(
			"FOR", pair.values,
				"expexted", pair.avarage,
				"got", v,
			)
		}
	}

}

func ExampleAvarage() {
	fmt.Println(Avarage([]float64{1, 2}))
	// Output:
	// 1.5
}
