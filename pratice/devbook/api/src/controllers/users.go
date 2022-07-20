package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
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
