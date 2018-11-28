// Package mjmath provides custom math solutions.
package mjmath

// Avarage calculates avarage of given numbers.
func Avarage(f []float64) float64 {
	if len(f) == 0 {
		return 0
	}
	var sum float64
	for _, v := range f {
		sum += v
	}
	r := sum / float64(len(f))
	return r
}

