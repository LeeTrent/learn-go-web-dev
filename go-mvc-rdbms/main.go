package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-rdbms/controller"
	"fmt"
	"log"
)

func main() {

	fmt.Println("main()")

	router := httprouter.New()
	ctrl := controller.NewBookController()

	// Index
	router.GET("/", ctrl.RetieveAllGet)

	// Retrieve
	router.GET("/books", ctrl.RetieveAllGet)
	router.GET("/books/retrieve/:isbn", ctrl.RetieveOneGet)

	// Create
	router.GET("/books/create", ctrl.CreateGet)
	router.POST("/books/create", ctrl.CreatePost)

	// Update
	router.GET("/books/update/:isbn", ctrl.UpdateGet)
	router.POST("/books/update", ctrl.UpdatePost)

	// Delete
	router.GET("/books/delete/:isbn", ctrl.DeleteGet)

	log.Fatal( http.ListenAndServe("localhost:8080", router) )
}