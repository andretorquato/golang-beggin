package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var userRequest models.User
	if erro = json.Unmarshal(bodyRequest, &userRequest); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userDB, erro := repository.FindByEmail(userRequest.Email)
	if erro != nil {
		response.Error(w, http.StatusNotFound, erro)
		return
	}

	if erro = security.CheckHash(userDB.Password, userRequest.Password); erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := authentication.GenerateToken(userDB.ID)

	userId := strconv.FormatUint(userDB.ID, 10)

	response.JSON(w, http.StatusOK, models.AuthData{ID: userId, Token: token})
}
