package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nick":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		log.Fatal(erro)
		return
	}

	response, erro := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if erro != nil {
		log.Fatal(erro)
		return
	}

	defer response.Body.Close()
	fmt.Println(bytes.NewBuffer(user))
	fmt.Println(response.Body)
}
