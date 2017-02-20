package main

import (
	"io"
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

		io.WriteString(conn, "I see you connected")
		conn.Close()
	}
}