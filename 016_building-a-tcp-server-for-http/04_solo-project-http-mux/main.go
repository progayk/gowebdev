package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err.Error())
	}

	i := 0
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err.Error())
		}
		// TODO: make it goroutine
		go handle(conn)
		i++
		log.Println( "connection is accepted.", i)
	}
}

func handle(conn net.Conn) {

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// send it to mux to be decided whether GET or POST
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// GET /selam HTTP/1.1
	m := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	switch m {
	case "GET":
		handleGet(conn, uri)
	case "POST":
		handlePost(conn, uri)
	default:
		fmt.Fprintln(conn, "ACCEPTED ONLY GET OR POST METHODS.")
	}
}

func handleGet(conn net.Conn, uri string) {
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
	// response
	respond(conn, page)
}

func handlePost(conn net.Conn, uri string) {
	page := uri + " Processed."
	respond(conn, page)
}

func respond(conn net.Conn, page string) {
	defer conn.Close()

	bs, err := ioutil.ReadFile("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	body := string(bs)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body, page)
}
