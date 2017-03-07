package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8080", nil)
}


func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("trent-cookie")

	if err != nil {
		if err == http.ErrNoCookie {
			cookie = &http.Cookie {
				Name: "trent-cookie",
				Value: "0",
			}
		} else {
			log.Fatal(err)
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}

	count++;
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}