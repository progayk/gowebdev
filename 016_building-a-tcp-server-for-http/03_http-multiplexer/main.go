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

// handleGet handles the GET method
func handleGet(uri string, c chan<- string) {
	var page string
	switch uri {
	case "/":
		page = "index"
	case "/about":
		page = "about"
	case "/apply":
		page = "apply"
	default:
		page = "404"
	}
	go func() {
		c <- page
	}()
	fmt.Println("GET method handled")
	fmt.Println("requested uri:", uri)
}

func handlePost(uri string,c chan<- string) {
	r := "Apply processing"
	go func() {
		c <- r
	}()
}

func mux(c chan<- string, reqLine string) {
	// request method
	bs := strings.Fields(reqLine)
	m := bs[0]
	uri := bs[1]
	switch m {
	case "GET":
		handleGet(uri, c)
	case "POST":
		handlePost(uri, c)
	default:
		fmt.Println("GET method ACCEPTED ONLY")
	}
	fmt.Println("***METHOD", m)
}

func request(conn net.Conn, c chan<- string) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			 mux(c, ln)
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
	body :=  fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>My own server build on top of a TCP server</title>
		</head>
		<body>
			<h1>%s page</h1>
			<a href="/">index</a><br>
			<a href="/about">about</a><br>
			<a href="/apply">apply</a><br><br>
			<form action="/apply" method="POST">
				<button type="submit">Apply</button>
			</form>
		</body>
	</html>`, v)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
