package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
			log.Println(err)
		}

		go handler(conn)

	}
}

func handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		t := time.Now()
		fmt.Fprint(conn,  t.Format("15:04"), " You\n\n")
		ln := scanner.Text()
		fmt.Fprintf(conn, "\t\t I heard you said %s <", ln)
		fmt.Fprintln(conn, "\n\t\t", t.Format("15:04"), " Mayk\n")
	}
	defer conn.Close()

	fmt.Println("how does reader know that the end of file is reached?")

}
