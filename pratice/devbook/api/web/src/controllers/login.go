package controllers

import "net/http"

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login screen"))
}
