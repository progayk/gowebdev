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
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Couldnt set the deadline")
	}

	go timer(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		t := time.Now()
		fmt.Fprint(conn,  t.Format("15:04"), " You\n\n")
		ln := scanner.Text()
		fmt.Fprintf(conn, "\t\t I heard you said %s <", ln)
		fmt.Fprintln(conn, "\n\t\t", t.Format("15:04"), " Mayk\n")
	}
	defer conn.Close()

	// Now it will know when we close
	fmt.Println("Now it will know when we close")

}

// timer informs the client about the time left to close the connection.
func timer(conn net.Conn) {
	for i := 10; i > 0; i-- {
		if i <= 3 || i == 10 {
			fmt.Fprintf(conn, "\n%d second(s) left to close the connection\n", i)
		}
		time.Sleep(1 * time.Second)
	}
}
