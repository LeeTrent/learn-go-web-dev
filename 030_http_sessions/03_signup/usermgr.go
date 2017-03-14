package main

import (
	"errors"
	"fmt"
)

type User struct {
	UserName string
	Password string
	First    string
	Last     string
}
type UserMgr map[string]User

func (um UserMgr) getUser(un string) User {
	return um[un]
}

func (um UserMgr) userNameIsTaken(un string) bool {
	_, ok := um[un]
	return ok
}

func (um UserMgr) removeUser(un string) {
	delete(um, un)
}

func (um UserMgr) createUser(un, pw, fn, ln string) (User, error) {
	if _, ok := um[un]; ok {
		errMsg := fmt.Sprintf("Username '%s' already taken", un)
		return User{}, errors.New(errMsg)
	}
	um[un] = User{un, pw, fn, ln}
	return um[un], nil
}
