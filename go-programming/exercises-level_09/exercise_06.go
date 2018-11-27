package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("System Info")
	fmt.Println("-----------")
	fmt.Println("the host OS:\t", runtime.GOOS)
	fmt.Println("the host ARCH:\t", runtime.GOARCH)
	fmt.Println("-----------")
}
