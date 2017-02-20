package main

import (
	"log"
	"net/http"
	"html/template"
)

var tmpl *template.Template

type content struct {
	Title string
	Data string
}

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog",  http.HandlerFunc(dog))
	http.Handle("/me",  http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	indexContent := content {
		Title: "index",
		Data: "index",
	}

	err := tmpl.ExecuteTemplate(res, "index.gohtml", indexContent)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {

	cnt := content {
		Title: "dog",
		Data: "dog",
	}

	err := tmpl.ExecuteTemplate(res, "index.gohtml", cnt)
	if err != nil {
		log.Fatalln(err)
	}
}

func me(res http.ResponseWriter, req *http.Request) {

	cnt := content {
		Title: "Lee Thomas Trent",
		Data: "Lee Thomas Trent",
	}

	err := tmpl.ExecuteTemplate(res, "index.gohtml", 	cnt )

	if err != nil {
		log.Fatalln(err)
	}
}