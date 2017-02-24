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


	var responseBody string

	switch {
	case method == "GET" && uri == "/":
		responseBody = bodyForGetIndex()
	case method == "GET" && uri == "/apply":
		responseBody = bodyForGetApply()
	case method == "POST" && uri == "/apply":
		responseBody = bodyForPostApply()
	default:
		responseBody = bodyForDefault()
	}


	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(responseBody))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, responseBody)
}


func bodyForGetIndex() string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
	`
}

func bodyForGetApply() string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET DOG</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
}

func bodyForPostApply() string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>
	`
}

func bodyForDefault() string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>DEFAULT</title>
		</head>
		<body>
			<h1>DEFAULT</h1>
		</body>
		</html>
	`
}