package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "someone1234"

	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("this is the slice of string", len(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte("someone123"))
	if err != nil {
		fmt.Println("You couldn't log in successfully")
		return
	}
	fmt.Println("You logged in succesfully.")
}
