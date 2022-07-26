package controllers

import (
	"api/src/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all Users"))
}

func GetAnUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get an User"))
}

func UpdateUserData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User data"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
