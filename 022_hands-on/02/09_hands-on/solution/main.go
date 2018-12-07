package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*"))
}

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
	var reqMethod, reqUri, msg string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			reql := strings.Fields(ln)
			reqMethod = reql[0]
			reqUri = reql[1]

			switch reqMethod {
			case "GET":
				if reqUri == "/" {
					msg = "index"
				}
				if reqUri == "/apply" {
					msg = "apply"
				}
			case "POST":
				if reqUri == "/apply" {
					msg = "applyProcess"
				}
			default:
				msg = "not a valid request"
			}

		}
		if ln == "" {
			fmt.Println("End of respond")
			break
		}
		fmt.Println(ln)
		i++
	}

	body := "Yes it is!"
	body += "\r\n"
	body += reqMethod
	body += "\r\n"
	body += reqUri
	body += "\n" + msg

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	err := tpl.ExecuteTemplate(conn, "index.html", body)
	if err != nil {
		log.Panic(err)
	}
}
