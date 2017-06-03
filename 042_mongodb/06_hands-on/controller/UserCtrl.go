package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/042_mongodb/06_hands-on/model"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/042_mongodb/06_hands-on/dao"
)

type UserController struct{
	userDao *dao.UserDAO
}

func NewUserController() *UserController {
	return &UserController{userDao: dao.NewUserDAO()}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Init User
	u := model.User{}

	// Populate User from Request Body
	json.NewDecoder(r.Body).Decode(&u)

	// Persist User
	u = uc.userDao.Create(u)

	// Marshal User model object to JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Prepare HTTP Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// Write response
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) RetrieveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Grab id from Request
	id := p.ByName("id")

	// Retrieve user from data store
	u, found := uc.userDao.Retrieve(id)

	if found {

		// Marshal User model object to JSON
		uj, err := json.Marshal(u)
		if err != nil {
			fmt.Println(err)
		}

		// Prepare HTTP Header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200

		// Write response
		fmt.Fprintf(w, "%s\n", uj)

	} else {

		// Prepare HTTP Header
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK) // 200

		// Write response
		fmt.Fprintf(w, "User with ID '%s' not found\n", id)
	}
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id from Request
	id := p.ByName("id")

	// Delete user from data store
	found := uc.userDao.Delete(id)

	// Prepare HTTP Header
	w.WriteHeader(http.StatusOK) // 200

	// Write response
	if found {
		fmt.Fprint(w, "Deleted user", id, "\n")
	} else {
		fmt.Fprintf(w, "User with ID '%s' not found\n", id)
	}
}

//func (uc UserController) RetrieveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// Grab id from Request
//	id := p.ByName("id")
//
//	// Retrieve user from data store
//	u := uc.userDao.Retrieve(id)
//
//	// Marshal User model object to JSON
//	uj, err := json.Marshal(u)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// Prepare HTTP Header
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK) // 200
//
//	// Write response
//	fmt.Fprintf(w, "%s\n", uj)
//}

//func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// Grab id from Request
//	id := p.ByName("id")
//
//	// Delete user from data store
//	uc.userDao.Delete(id)
//
//	// Prepare HTTP Header
//	w.WriteHeader(http.StatusOK) // 200
//
//	// Write response
//	fmt.Fprint(w, "Deleted user", id, "\n")
//}