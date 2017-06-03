package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/042_mongodb/06_hands-on/controller"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controller.NewUserController()
	r.GET("/user/:id", uc.RetrieveUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}