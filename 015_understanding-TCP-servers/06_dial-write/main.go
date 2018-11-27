package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprintln(conn, "I dialed you.")

	defer conn.Close()

}

