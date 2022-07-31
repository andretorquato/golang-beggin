package server

import (
	"database-pratice/src/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"username,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user user

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO users (username, nickname, email, password) VALUES (?, ?, ?, ?)")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	inserted, erro := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	insertedID, erro := inserted.LastInsertId()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"id\": %d}", insertedID)))
}
