package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = user.Prepare("create"); erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	repository := repositories.NewUsersRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	queryUser := strings.ToLower(r.URL.Query().Get("user"))
	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, erro := repository.Search(queryUser)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func GetAnUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	userIDIntoToken, erro := authentication.GetUserID(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDIntoToken {
		response.Error(w, http.StatusForbidden, errors.New("unauthorized, you can only get your own data"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, erro := repository.FindByID(userID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUserData(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	userIDIntoToken, erro := authentication.GetUserID(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDIntoToken {
		response.Error(w, http.StatusForbidden, errors.New("unauthorized, you can only update your own data"))
		return
	}

	var user models.User
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("update"); erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	erro = repository.Update(user, userID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	userIDIntoToken, erro := authentication.GetUserID(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDIntoToken {
		response.Error(w, http.StatusForbidden, errors.New("unauthorized, you can only delete your own data"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	erro = repository.Delete(userID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)

}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	userIDIntoToken, erro := authentication.GetUserID(r)

	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	followUser, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if followUser == userIDIntoToken {
		response.Error(w, http.StatusBadRequest, errors.New("you can't follow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	erro = repository.Follow(userIDIntoToken, followUser)

	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	userIDIntoToken, erro := authentication.GetUserID(r)

	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	unFollowerID, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if unFollowerID == userIDIntoToken {
		response.Error(w, http.StatusBadRequest, errors.New("you can't unfollow yourself"))
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	erro = repository.UnFollow(userIDIntoToken, unFollowerID)

	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewUsersRepository(db)
	followers, erro := repository.GetFollowers(userID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewUsersRepository(db)
	following, erro := repository.GetFollowing(userID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, following)
}
