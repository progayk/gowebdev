package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMaykMath = errors.New("mayk math: square root of negative number")

func main() {
	fmt.Printf("Type of ErrMaykMath is %T\n", ErrMaykMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
	
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMaykMath
	}
	return 42, nil
}