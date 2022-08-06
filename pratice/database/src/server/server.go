package server

import (
	"database-pratice/src/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connect()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, erro := db.Query("SELECT * FROM users")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		erro = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.Atoi(params["id"])
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row, erro := db.Query("SELECT * FROM users WHERE id = ?", ID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var user user
	for row.Next() {
		erro = row.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.Atoi(params["id"])
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	statement, erro := db.Prepare("UPDATE users SET username = ?, nickname = ?, email = ?, password = ? WHERE id = ?")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, erro = statement.Exec(user.Name, user.Nickname, user.Email, user.Password, ID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"User updated\"}"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseInt(params["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, erro = statement.Exec(ID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"User deleted\"}"))
}
