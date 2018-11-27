package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := "8080"
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	log.Printf("Server is running on localhost port %v\n", port)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Println("============")
		log.Println("a request is recieved")
		log.Println("============")

		io.WriteString(conn, "\nHello from TCP server.\n")
		fmt.Fprintln(conn, "how is your day")
		fmt.Fprintf(conn, "%v", "Well I Hope\n\n")

		time.Sleep(10000*time.Millisecond)
		conn.Close()
	}

}
