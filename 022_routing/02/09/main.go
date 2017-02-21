package main

import (
	"io"
	"bufio"
	"fmt"
	"log"
	"net"
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

	fmt.Println(conn, "I see you connected.\n")
	io.WriteString(conn, "I see you connected.\n")

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		if "" == line {
			break
		}
		fmt.Println(line)
		fmt.Fprintf(conn, "fmt.Fprintf(): %s\n", line)

	}

	fmt.Println("Broke out of Scanner for loop.")
	io.WriteString(conn, "fmt.Fprintf(): Broke out of Scanner for loop.")

}