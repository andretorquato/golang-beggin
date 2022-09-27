package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/models"
	"web/src/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(user))

	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CaptureStatusCodeError(w, response)
		return
	}

	var authData models.AuthData
	if erro = json.NewDecoder(response.Body).Decode(&authData); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println(authData)
	responses.JSON(w, response.StatusCode, nil)
}
