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
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	// create a in-memory database as a map
	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n\n", v)
			//fmt.Printf("The user requested the value %s with key %s\n\n", v, k)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
			//fmt.Printf("The user set a new key as %s and a value as %s\n\n", k, v)
		case "DEL":
			k := fs[1]
			delete(data, k)
			//fmt.Printf("The user deleted the key %s\n\n", k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}
	}
}