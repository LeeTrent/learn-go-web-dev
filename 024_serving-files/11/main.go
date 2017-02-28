package main

import (
	"html/template"
	"net/http"
	"log"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)

	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, _ *http.Request) {
	handleResponse(resp, "index.gohtml")
}

func about(resp http.ResponseWriter, _ *http.Request) {
	handleResponse(resp, "about.gohtml")
}

func contact(resp http.ResponseWriter, _ *http.Request) {
	handleResponse(resp, "contact.gohtml")
}

func apply(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		handleResponse(resp, "applyProcess.gohtml")
	} else {
		handleResponse(resp, "apply.gohtml")
	}
}

func handleResponse(resp http.ResponseWriter, templateName string) {
	err := templ.ExecuteTemplate(resp, templateName, nil)
	handleErrorIfAny(resp, err)
}

func handleErrorIfAny(resp http.ResponseWriter, err error) {
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}