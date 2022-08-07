package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
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

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewUsersRepository(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id added %d", userID)))
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
