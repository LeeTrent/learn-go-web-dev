package main

import (
	"io"
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	list, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer list.Close()

	for {
		conn, err := list.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	var method, uri string

	counter := 0
	for scanner.Scan() {
		line := scanner.Text()

		if counter == 0 {
			method = strings.Fields(line)[0]
			uri = strings.Fields(line)[1]

			fmt.Println("METHOD: ", method)
			fmt.Println("URI: ", uri)

		}
		if "" == line {
			fmt.Println("(fmt.Println): End of Request reached.")
			break
		}

		counter++
	}

	responseBody := fmt.Sprintf("Method is '%s' and URI is '%s'\r\n", method, uri)

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(responseBody))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, responseBody)
}