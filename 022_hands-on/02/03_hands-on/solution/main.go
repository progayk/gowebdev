package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			fmt.Println("End of respond")
			break
		}
		fmt.Println(ln)
	}

	fmt.Printf("selam dostum: %s\n", "mayk")
	io.WriteString(conn, "i see you connected.")
}

//func myLog(msg string) {
//	f, err := os.OpenFile("log.txt", os.O_RDWR, 0755)
//	if err != nil {
//		log.Println("couldnt open file: ", err)
//	}
//	defer f.Close()
//
//	_, err = io.WriteString(f, msg)
//	if err != nil {
//		log.Fatalln("couldnt copy to file: ", err)
//	}
//}
