package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
	
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number: %v", f)
	}
	return 42, nil
}