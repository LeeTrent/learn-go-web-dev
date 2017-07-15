package view

import (
	"html/template"
	"net/http"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-rdbms/model"
	"fmt"
)

const (
	errorTemplate     		= "error.gohtml"
	createGetTemplate 		= "createGet.gohtml"
	createPostTemplate		= "createPost.gohtml"
	updateGetTemplate   	= "updateGet.gohtml"
	updatePostTemplate  	= "updatePost.gohtml"
	retrieveAllGetTemplete 	= "retrieveAll.gohtml"
	retrieveOneGetTemplete 	= "retrieveOne.gohtml"
)

type BookView struct {
	templ *template.Template
}

func NewBookView() *BookView {

	fmt.Println("NewBookView()")

	t := template.Must(template.ParseGlob("view/templates/*.gohtml"))
	return &BookView{templ: t}
}

func (bv *BookView) CreateGet(resp http.ResponseWriter) {
	bv.templ.ExecuteTemplate(resp, createGetTemplate, nil)
}

func (bv *BookView) CreatePost(resp http.ResponseWriter, book model.Book) {
	bv.templ.ExecuteTemplate(resp, createPostTemplate, book)
}

func (bv *BookView) UpdateGet(resp http.ResponseWriter, book model.Book) {
	bv.templ.ExecuteTemplate(resp, updateGetTemplate, book)
}

func (bv *BookView) UpdatePost(resp http.ResponseWriter, book model.Book) {
	bv.templ.ExecuteTemplate(resp, updatePostTemplate, book)
}

func (bv *BookView) RetrieveAllGet(resp http.ResponseWriter, books []model.Book) {
	bv.templ.ExecuteTemplate(resp, retrieveAllGetTemplete, books)
}

func (bv *BookView) RetrieveOneGet(resp http.ResponseWriter, book model.Book) {
	bv.templ.ExecuteTemplate(resp, retrieveOneGetTemplete, book)
}

func (bv *BookView) Error(resp http.ResponseWriter, errMsg string ) {
	bv.templ.ExecuteTemplate(resp, errorTemplate, errMsg)
}