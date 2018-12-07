package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}
}


func handleConn(conn net.Conn) {
	defer conn.Close()

	i := 0
	var reqMethod, reqURI string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			reql := strings.Fields(ln)
			reqMethod = reql[0]
			reqURI = reql[1]
			fmt.Println("method: ", reqMethod)
			fmt.Println("uri: ", reqURI)
		}
		if ln == "" {
			fmt.Println("End of respond")
			break
		}
		fmt.Println(ln)
		i++
	}

	body := "Yes it is!"

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	//fmt.Fprint(conn, "Mayk-Type: hahahah")
	io.WriteString(conn, "\r\n")
	body += "\r\n"
	body += reqMethod
	body += "\r\n"
	body += reqURI
	io.WriteString(conn, body)
}
