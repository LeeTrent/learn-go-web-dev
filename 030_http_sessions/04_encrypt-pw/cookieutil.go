package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func CreateAndSetCookie(cookieName, cookieValue string, response http.ResponseWriter) *http.Cookie {
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	}
	http.SetCookie(response, cookie)
	return cookie
}

func GetAndSetCookie(cookieName string, w http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie(cookieName)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  cookieName,
			Value: uuid.NewV4().String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func HasCookie(cookieName string, req *http.Request) bool {
	_, err := req.Cookie(cookieName)
	if err != nil && err == http.ErrNoCookie {
		return false
	}
	return true
}
