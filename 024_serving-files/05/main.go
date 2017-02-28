package main

import (
	"html/template"
	"log"
	"net/http"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}


func index(res http.ResponseWriter, req *http.Request ) {
	err := templ.Execute(res, nil)
	if err != nil {
		log.Fatalln("Template 'templates/index.gohtml' did not execute", err)
	}
}