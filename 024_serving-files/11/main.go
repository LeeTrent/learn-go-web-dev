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
	handleResponse(resp, "index.gohtml", nil)
}

func about(resp http.ResponseWriter, _ *http.Request) {
	handleResponse(resp, "about.gohtml", nil)
}

func contact(resp http.ResponseWriter, _ *http.Request) {
	handleResponse(resp, "contact.gohtml", nil)
}

func apply(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		handleResponse(resp, "applyProcess.gohtml", nil)
	} else {
		handleResponse(resp, "apply.gohtml", nil)
	}
}

func handleResponse(resp http.ResponseWriter, templateName string, data interface{}) {
	err := templ.ExecuteTemplate(resp, templateName, data)
	handleErrorIfAny(resp, err)
}

func handleErrorIfAny(resp http.ResponseWriter, err error) {
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}