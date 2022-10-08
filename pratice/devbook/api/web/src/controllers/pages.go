package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
	"web/src/utils"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadRegisterScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoadHomeScreen(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	res, erro := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	var posts []models.Post

	if erro = json.NewDecoder(res.Body).Decode(&posts); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "home.html", posts)
}
