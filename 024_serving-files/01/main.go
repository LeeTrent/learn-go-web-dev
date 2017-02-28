package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var dogTemplate *template.Template

func init() {
	dogTemplate = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPic)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func foo(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "foo ran")
}

func dog(resp http.ResponseWriter, req *http.Request) {
	dogTemplate.ExecuteTemplate(resp, "dog.gohtml", nil)
}

func dogPic(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, "dog.jpg")
}