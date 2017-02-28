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
	http.HandleFunc("/", index)
	http.Handle("/resources/",
		http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}


func index(res http.ResponseWriter, req *http.Request ) {
	err := templ.Execute(res, nil)
	if err != nil {
		log.Fatalln("Template 'templates/index.gohtml' did not execute", err)
	}
}