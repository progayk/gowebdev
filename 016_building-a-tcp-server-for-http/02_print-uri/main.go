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
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()

	c := make(chan string)

	request(conn, c)
	response(conn, c)
}

func request(conn net.Conn, c chan<- string) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request method
			bs := strings.Fields(ln)
			m := bs[0]
			uri := bs[1]
			go func() {
				c <- uri
			}()
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			// headers are done
			fmt.Println("\n\nheaders are done.\n\n")
			break
		}
		i++
	}
}

func response(conn net.Conn, c <-chan string) {
	v := <-c
	body :=  fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>%s</strong></body></html>`, v)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
