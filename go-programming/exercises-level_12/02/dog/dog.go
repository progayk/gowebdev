// Package dog provides dog related calculations.
package dog

// ConvToDogYears takes an human age and converts it into dog years.
// 7 human years is equal to 1 dog year.
func ConvToDogYears(n int) float64 {
	dogYear := float64(7)
	return float64(n) / dogYear
}

