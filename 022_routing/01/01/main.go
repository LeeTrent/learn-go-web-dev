package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "index")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Lee Trent")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

