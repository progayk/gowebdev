package mjmath

import "testing"

func TestAvarage(t *testing.T) {
	var v float64
	v = Avarage([]float64{1, 2})
	if v != 1.5 {
		t.Errorf("Expected 1.5 got %v", v)
	}

}
