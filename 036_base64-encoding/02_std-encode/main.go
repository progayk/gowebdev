package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "I'm still breathing, I'm still breathing. I'm aliiiiiiiive. I'm aliiiiiiive..."

	encoded := encode(s)
	fmt.Println(encoded)

	decoded := decode(encoded)
	fmt.Println(string(decoded))

	fmt.Println(encoded == base64.StdEncoding.EncodeToString([]byte(decoded)))
}

func encode(s string) string {
	e := base64.StdEncoding.EncodeToString([]byte(s))
	return e
}

func decode(s string) string {
	e, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Println("error decoding", err)
	}
	return string(e)
}
