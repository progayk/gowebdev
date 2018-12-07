// Package acdc asks if you are ready to rock.
package acdc

// Sum adds any number of value of type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
