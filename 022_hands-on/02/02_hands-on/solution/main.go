package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}
}


func handleConn(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		go myLog(ln)
	}

	fmt.Fprintf(conn, "selam dostum: %s\n", "mayk")
	io.WriteString(conn, "i see you connected.")
}

func myLog(msg string) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("couldnt open file: ", err)
	}
	defer f.Close()

	_, err = io.WriteString(f, msg)
	if err != nil {
		log.Fatalln("couldnt copy to file: ", err)
	}
}
