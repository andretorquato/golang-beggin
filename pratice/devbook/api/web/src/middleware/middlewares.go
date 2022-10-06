package middleware

import (
	"fmt"
	"log"
	"net/http"
	"web/src/cookies"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values, erro := cookies.Read(r)
		if erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		fmt.Println(values, erro)
		nextFunc(w, r)
	}
}
