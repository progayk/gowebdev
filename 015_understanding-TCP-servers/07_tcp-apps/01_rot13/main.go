package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("couldnt accept conn", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
        scanner := bufio.NewScanner(conn)

        for scanner.Scan() {
        	ln := strings.ToLower(scanner.Text())
        	bs := []byte(ln)
        	r13 := rot13(bs)
        	fmt.Fprintf(conn, "%s - %s\n\n", ln, r13)
        }
}

func rot13(msg []byte) []byte {
	bs := make([]byte, len(msg))
	for i, v := range msg {
		// ascii 97 - 122
		if msg[i] <= 109 {
			bs[i] = v + 13
		} else {
			bs[i] = v - 13
		}
	}
	return bs
}
