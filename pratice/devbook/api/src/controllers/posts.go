package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.GetUserID(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro = json.Unmarshal(bodyRequest, &post); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Prepare(); erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	post.AuthorID = userID
	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewPostsRepository(db)
	post.ID, erro = repository.Create(post)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	userIDIntoToken, erro := authentication.GetUserID(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	userID := uint64(userIDIntoToken)
	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, erro := repository.GetAll(userID)

	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func GetAnPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post, erro := repository.GetByID(postID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletePost"))
}
