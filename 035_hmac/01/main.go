package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	c := genHMAC("test@test.com", "privatekey")
	fmt.Println(c)
	c = genHMAC("test@test.com", "privatekey")
	fmt.Println(c)
}

func genHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return fmt.Sprintf("%x", mac.Sum(nil))
}