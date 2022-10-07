package controllers

import (
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/requests"
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
	response, erro := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, erro)

	utils.ExecuteTemplate(w, "home.html", nil)
}
